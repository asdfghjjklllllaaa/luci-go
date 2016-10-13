// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package delegation

import (
	"encoding/base64"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"

	"github.com/luci/luci-go/common/clock"
	"github.com/luci/luci-go/common/clock/testclock"
	"github.com/luci/luci-go/common/logging"
	"github.com/luci/luci-go/common/logging/memlogger"
	"github.com/luci/luci-go/server/auth/identity"
	"github.com/luci/luci-go/server/auth/signing"
	"github.com/luci/luci-go/server/auth/signing/signingtest"

	"github.com/luci/luci-go/server/auth/delegation/messages"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckToken(t *testing.T) {
	c := memlogger.Use(context.Background())
	c, _ = testclock.UseTime(c, testclock.TestRecentTimeUTC)
	minter := newFakeTokenMinter()

	defer func() {
		if t.Failed() {
			logging.Get(c).(*memlogger.MemLogger).Dump(os.Stderr)
		}
	}()

	Convey("Basic use case", t, func() {
		tok := minter.mintToken(c, subtoken(c, "user:from@example.com", "user:to@example.com"))
		ident, err := CheckToken(c, CheckTokenParams{
			Token:                tok,
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldBeNil)
		So(ident, ShouldEqual, "user:from@example.com")
	})

	Convey("Basic use case with group check", t, func() {
		tok := minter.mintToken(c, subtoken(c, "user:from@example.com", "group:token-users"))

		groups := &fakeGroups{
			groups: map[string]string{
				"token-users": "user:to@example.com",
			},
		}

		// Pass.
		ident, err := CheckToken(c, CheckTokenParams{
			Token:                tok,
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        groups,
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldBeNil)
		So(ident, ShouldEqual, "user:from@example.com")

		// Fail.
		_, err = CheckToken(c, CheckTokenParams{
			Token:                tok,
			PeerID:               "user:NOT-to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        groups,
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldEqual, ErrForbiddenDelegationToken)
	})

	Convey("Not base64", t, func() {
		_, err := CheckToken(c, CheckTokenParams{
			Token:                "(^*#%^&#%",
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldEqual, ErrMalformedDelegationToken)
	})

	Convey("Huge token is skipped", t, func() {
		_, err := CheckToken(c, CheckTokenParams{
			Token:                strings.Repeat("aaaa", 10000),
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldEqual, ErrMalformedDelegationToken)
	})

	Convey("Untrusted signer", t, func() {
		tok := minter.mintToken(c, subtoken(c, "user:from@example.com", "user:to@example.com"))
		minter.signerID = "service:nah-i-renamed-myself"
		_, err := CheckToken(c, CheckTokenParams{
			Token:                tok,
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldEqual, ErrUnsignedDelegationToken)
	})

	Convey("Bad signature", t, func() {
		tok := minter.mintToken(c, subtoken(c, "user:from@example.com", "user:to@example.com"))
		// An offset in serialized token that points to Subtoken field. Replace one
		// byte there to "break" the signature.
		sigOffset := len(tok) - 10
		So(tok[sigOffset], ShouldNotEqual, 'A')
		_, err := CheckToken(c, CheckTokenParams{
			Token:                tok[:sigOffset] + "A" + tok[sigOffset+1:],
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldEqual, ErrUnsignedDelegationToken)
	})

	Convey("Expired token", t, func() {
		tok := minter.mintToken(c, subtoken(c, "user:from@example.com", "user:to@example.com"))

		clock.Get(c).(testclock.TestClock).Add(2 * time.Hour)

		_, err := CheckToken(c, CheckTokenParams{
			Token:                tok,
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldEqual, ErrForbiddenDelegationToken)
	})

	Convey("Wrong target service", t, func() {
		tok := minter.mintToken(c, subtoken(c, "user:from@example.com", "user:to@example.com"))
		_, err := CheckToken(c, CheckTokenParams{
			Token:                tok,
			PeerID:               "user:to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:NOT-a-service-id",
		})
		So(err, ShouldEqual, ErrForbiddenDelegationToken)
	})

	Convey("Wrong audience", t, func() {
		tok := minter.mintToken(c, subtoken(c, "user:from@example.com", "user:to@example.com"))
		_, err := CheckToken(c, CheckTokenParams{
			Token:                tok,
			PeerID:               "user:NOT-to@example.com",
			CertificatesProvider: minter,
			GroupsChecker:        &fakeGroups{},
			OwnServiceIdentity:   "service:service-id",
		})
		So(err, ShouldEqual, ErrForbiddenDelegationToken)
	})

}

// subtoken returns messages.Subtoken with some fields filled in.
func subtoken(c context.Context, delegatedID, audience string) *messages.Subtoken {
	return &messages.Subtoken{
		DelegatedIdentity: delegatedID,
		CreationTime:      clock.Now(c).Unix() - 300,
		ValidityDuration:  3600,
		Audience:          []string{audience},
		Services:          []string{"service:service-id"},
	}
}

// fakeTokenMinter knows how to generate tokens.
//
// It also implements CertificatesProvider protocol that is used when validating
// the tokens.
type fakeTokenMinter struct {
	signer   signing.Signer
	signerID string
}

func newFakeTokenMinter() *fakeTokenMinter {
	return &fakeTokenMinter{
		signer:   signingtest.NewSigner(0, nil),
		signerID: "service:fake-signer",
	}
}

func (f *fakeTokenMinter) GetCertificates(c context.Context, id identity.Identity) (*signing.PublicCertificates, error) {
	if string(id) != f.signerID {
		return nil, nil
	}
	return f.signer.Certificates(c)
}

func (f *fakeTokenMinter) mintToken(c context.Context, subtoken *messages.Subtoken) string {
	blob, err := proto.Marshal(subtoken)
	if err != nil {
		panic(err)
	}
	keyID, sig, err := f.signer.SignBytes(c, blob)
	if err != nil {
		panic(err)
	}
	tok, err := proto.Marshal(&messages.DelegationToken{
		SerializedSubtoken: blob,
		SignerId:           f.signerID,
		SigningKeyId:       keyID,
		Pkcs1Sha256Sig:     sig,
	})
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(tok)
}

// fakeGroups implements GroupsChecker.
type fakeGroups struct {
	groups map[string]string // if nil == no group checks
}

func (f *fakeGroups) IsMember(c context.Context, id identity.Identity, groups ...string) (bool, error) {
	if f.groups == nil {
		return true, nil
	}
	for _, group := range groups {
		if f.groups[group] == string(id) {
			return true, nil
		}
	}
	return false, nil
}
