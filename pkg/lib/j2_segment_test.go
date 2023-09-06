// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package lib

import (
	"bytes"
	"encoding/json"

	"gopkg.in/check.v1"

	"github.com/bloomcredit/moov-metro2/pkg/utils"
)

func (t *SegmentTest) TestJ2Segment(c *check.C) {
	segment := NewJ2Segment()
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(0, check.Equals, bytes.Compare(segment.Bytes(), t.sampleJ2Segment))
	c.Assert(segment.Name(), check.Equals, J2SegmentName)
	c.Assert(segment.Length(), check.Equals, J2SegmentLength)
}

func (t *SegmentTest) TestJ2SegmentWithInvalidData(c *check.C) {
	segment := NewJ2Segment()
	_, err := segment.Parse(append([]byte("ERROR"), t.sampleJ2Segment...))
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ2SegmentWithInvalidGenerationCode(c *check.C) {
	segment := J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.GenerationCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "generation code in j2 segment has an invalid value")
}

func (t *SegmentTest) TestJ2SegmentWithEmptyGenerationCode(c *check.C) {
	jsonStr := `{
      "segmentIdentifier": "J2",
      "surname": "BEAUCHAMP",
      "firstName": "KEVIN",
      "socialSecurityNumber": 445112877,
      "dateBirth": "2020-01-02T00:00:00Z",
      "telephoneNumber": 4335552333,
      "ecoaCode": "2",
      "consumerInformationIndicator": "R",
      "countryCode": "US",
      "firstLineAddress": "234 HARRISON PLACE",
      "secondLineAddress": "SUITE #305",
      "city": "LANSING",
      "state": "MI",
      "zipCode": "72654",
      "addressIndicator": "Y",
      "residenceCode": "O"
    }`
	segment := NewJ2Segment()
	err := json.Unmarshal([]byte(jsonStr), &segment)
	c.Assert(err, check.IsNil)
	err = segment.Validate()
	c.Assert(err, check.IsNil)
	c.Assert(segment.Name(), check.Equals, J2SegmentName)
	c.Assert(segment.Length(), check.Equals, J2SegmentLength)
}

func (t *SegmentTest) TestJ2SegmentWithInvalidTelephoneNumber(c *check.C) {
	segment := &J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.TelephoneNumber = 0
	err = segment.Validate()
	c.Assert(err, check.IsNil)
}

func (t *SegmentTest) TestJ2SegmentWithInvalidAddressIndicator(c *check.C) {
	segment := J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.AddressIndicator = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "address indicator in j2 segment has an invalid value")
}

func (t *SegmentTest) TestJ2SegmentWithInvalidResidenceCode(c *check.C) {
	segment := J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)
	segment.ResidenceCode = "0"
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
	c.Assert(err.Error(), check.DeepEquals, "residence code in j2 segment has an invalid value")
}

func (t *SegmentTest) TestJ2SegmentWithInvalidData2(c *check.C) {
	_, err := NewJ2Segment().Parse(t.sampleJ2Segment[:16])
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ2SegmentWithSocialSecurityNumber(c *check.C) {
	segment := &J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)

	segment.SocialSecurityNumber = 0
	err = segment.Validate()
	c.Assert(err, check.Equals, nil)

	segment.DateBirth = utils.Time{}
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
}

func (t *SegmentTest) TestJ2SegmentWithDateBirth(c *check.C) {
	segment := &J2Segment{}
	_, err := segment.Parse(t.sampleJ2Segment)
	c.Assert(err, check.IsNil)

	segment.DateBirth = utils.Time{}
	err = segment.Validate()
	c.Assert(err, check.Equals, nil)

	segment.SocialSecurityNumber = 0
	err = segment.Validate()
	c.Assert(err, check.Not(check.IsNil))
}
