// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package main

import . "gopkg.in/check.v1"

func (s *MySuite) TestExtractConceptWithoutParameters(c *C) {
	STEP := "* step that takes a table \"arg\""
	heading, text, conceptText, hasParam, _ := getTextForConcept(STEP)

	c.Assert(text, Equals, STEP+"\n")
	c.Assert(heading, Equals, "# "+CONCEPT_HEADING_TEMPLATE)
	c.Assert(conceptText, Equals, "* "+CONCEPT_HEADING_TEMPLATE)
	c.Assert(hasParam, Equals, false)
}

func (s *MySuite) TestExtractConcept(c *C) {
	STEP := "* step that takes a table \"arg\""
	heading, text, conceptText, hasParam, _ := getTextForConcept(STEP + "\n" + STEP)

	c.Assert(text, Equals, "* step that takes a table <arg>\n* step that takes a table <arg>\n")
	c.Assert(heading, Equals, "# "+CONCEPT_HEADING_TEMPLATE+" <arg>")
	c.Assert(conceptText, Equals, "* "+CONCEPT_HEADING_TEMPLATE+" \"arg\"")
	c.Assert(hasParam, Equals, true)
}

func (s *MySuite) TestExtractConceptWithTableAsArg(c *C) {
	STEP := `* Step that takes a table
	|Product|Description                  |
	|-------|-----------------------------|
	|Gauge  |BDD style testing with ease  |
	|Mingle |Agile project management     |
	|Snap   |Hosted continuous integration|
	|Gocd   |Continuous delivery platform |
	* Step that takes a table
	|Product|Description                  |
	|-------|-----------------------------|
	|Gauge  |BDD style testing with ease  |
	|Mingle |Agile project management     |
	|Snap   |Hosted continuous integration|
	|Gocd   |Continuous delivery platform |
	`
	heading, text, conceptText, hasParam, _ := getTextForConcept(STEP)

	c.Assert(text, Equals, "* Step that takes a table <table>\n* Step that takes a table <table>\n")
	c.Assert(heading, Equals, "# "+CONCEPT_HEADING_TEMPLATE+" <table>")
	c.Assert(conceptText, Equals, "* "+CONCEPT_HEADING_TEMPLATE+`
     |Product|Description                  |
     |-------|-----------------------------|
     |Gauge  |BDD style testing with ease  |
     |Mingle |Agile project management     |
     |Snap   |Hosted continuous integration|
     |Gocd   |Continuous delivery platform |
`)
	c.Assert(hasParam, Equals, true)
}

func (s *MySuite) TestExtractConceptWithTablesAsArg(c *C) {
	STEP := `* Step that takes a table
	|Product|Description                  |
	|-------|-----------------------------|
	|Gauge  |BDD style testing with ease  |
	|Mingle |Agile project management     |
	|Snap   |Hosted continuous integration|
	|Gocd   |Continuous delivery platform |
	* Step that takes a table
	|Product|Description                  |
	|-------|-----------------------------|
	|Gauge  |BDD style testing with ease  |
	* Step that takes a table
	|Product|Description                  |
	|-------|-----------------------------|
	|Gauge  |BDD style testing with ease  |
	|Mingle |Agile project management     |
	|Snap   |Hosted continuous integration|
	|Gocd   |Continuous delivery platform |
	`
	heading, text, conceptText, hasParam, _ := getTextForConcept(STEP)

	c.Assert(text, Equals, `* Step that takes a table <table>
* Step that takes a table 
     |Product|Description                |
     |-------|---------------------------|
     |Gauge  |BDD style testing with ease|
* Step that takes a table <table>
`)
	c.Assert(heading, Equals, "# "+CONCEPT_HEADING_TEMPLATE+" <table>")
	c.Assert(conceptText, Equals, "* "+CONCEPT_HEADING_TEMPLATE+`
     |Product|Description                  |
     |-------|-----------------------------|
     |Gauge  |BDD style testing with ease  |
     |Mingle |Agile project management     |
     |Snap   |Hosted continuous integration|
     |Gocd   |Continuous delivery platform |
`)
	c.Assert(hasParam, Equals, true)
}
