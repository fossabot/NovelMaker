/*
@Author: Hughie
@CreateTime: 2024-8-2
@LastEditors: Hughie
@LastEditTime: 2024-08-7
@Description: The definition of the xml Format
*/

package epub

import "encoding/xml"

/*
*
* content.opt
*
 */

type MetaNode struct {
	XMLName xml.Name `xml:"meta"`
	Name    string   `xml:"name,attr,omitempty"`
	Content string   `xml:"content,attr"`
	Equiv   string   `xml:"http-equiv,attr,omitempty"`
}

type DCCreator struct {
	XMLName xml.Name `xml:"dc:creator"`
	Id      string   `xml:"id,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type DCContributor struct {
	XMLName xml.Name `xml:"dc:contributor"`
	Id      string   `xml:"id,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type DCSubject struct {
	XMLName xml.Name `xml:"dc:subject"`
	Id      string   `xml:"id,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Identifier struct {
	XMLName xml.Name `xml:"dc:identifier"`
	Id      string   `xml:"id,attr"`
	Value   string   `xml:",chardata"`
}

type MetadataNode struct {
	XMLName      xml.Name        `xml:"metadata"`
	Xmlns        string          `xml:"xmlns:dc,attr"`
	Opt          string          `xml:"xmlns:opf,attr"`
	Title        string          `xml:"dc:title"`
	Creators     []DCCreator     `xml:"dc:creator"`
	Identifier   Identifier      `xml:"dc:identifier"`
	Language     string          `xml:"dc:language"`
	Contributors []DCContributor `xml:"dc:contributor"`
	Description  string          `xml:"dc:description"`
	Publisher    string          `xml:"dc:publisher"`
	Subject      []DCSubject     `xml:"dc:subject"`
	Date         string          `xml:"dc:date"`
	Metas        []MetaNode      `xml:"meta"`
}

type ItemNode struct {
	XMLName    xml.Name `xml:"item"`
	Id         string   `xml:"id,attr"`
	Href       string   `xml:"href,attr"`
	MediaType  string   `xml:"media-type,attr"`
	Properties string   `xml:"properties,attr,omitempty"`
}

type ManifestNode struct {
	XMLName xml.Name   `xml:"manifest"`
	Items   []ItemNode `xml:"item"`
}

type SpineItemNode struct {
	XMLName xml.Name `xml:"itemref"`
	Idref   string   `xml:"idref,attr"`
}

type SpineNode struct {
	XMLName xml.Name        `xml:"spine"`
	Toc     string          `xml:"toc,attr"`
	Items   []SpineItemNode `xml:"itemref"`
}

type PackageNode struct {
	XMLName    xml.Name     `xml:"package"`
	Xmlns      string       `xml:"xmlns,attr"`
	Identifier string       `xml:"unique-identifier,attr"`
	Version    string       `xml:"version,attr"`
	Metadata   MetadataNode `xml:"metadata"`
	Manifest   ManifestNode `xml:"manifest"`
	Spine      SpineNode    `xml:"spine"`
}

/*
*
* toc.ncx
*
 */

type HeadNode struct {
	XMLName xml.Name   `xml:"head"`
	Meta    []MetaNode `xml:"meta"`
}

type TextNode struct {
	Text string `xml:"text"`
}

type Content struct {
	Src string `xml:"src,attr"`
}

type NavPoint struct {
	Navlable  TextNode   `xml:"navLabel"`
	Content   Content    `xml:"content"`
	Id        string     `xml:"id,attr"`
	PlayOrder int        `xml:"playOrder,attr"`
	NavPoints []NavPoint `xml:"navPoint"`
}

type NavMap struct {
	NavPoints []NavPoint `xml:"navPoint"`
}

type NcxNode struct {
	XMLName xml.Name   `xml:"ncx"`
	Xmlns   string     `xml:"xmlns,attr"`
	Version string     `xml:"version,attr"`
	Header  HeadNode   `xml:"head"`
	Title   TextNode   `xml:"docTitle"`
	Author  []TextNode `xml:"docAuthor"`
	NavMap  NavMap     `xml:"navMap"`
}

/*
*
* nav.xhtml
*
 */
type Link struct {
	XMLName xml.Name `xml:"link"`
	Href    string   `xml:"href,attr"`
	Rel     string   `xml:"rel,attr"`
	Type    string   `xml:"type,attr"`
}

type HtmlHead struct {
	Title string `xml:"title"`
	Link  []Link `xml:"link"`
}

type NaVA struct {
	Text string `xml:",chardata"`
	Href string `xml:"href,attr"`
}

type NavDiv struct {
	XMLName  xml.Name `xml:"div"`
	Class    string   `xml:"class,attr"`
	Navlable NaVA     `xml:"a"`
	Child    []NavDiv `xml:"div"`
}

type NavBody struct {
	XMLName xml.Name `xml:"body"`
	Nav     []NavDiv `xml:"nav"`
}

type NavHTML struct {
	XMLName xml.Name `xml:"html"`
	Xmlns   string   `xml:"xmlns,attr"`
	Header  HtmlHead `xml:"head"`
	Body    NavBody  `xml:"body"`
}

/*
*
* text.xhtml
*
 */
type XhtmlBody struct {
	XMLName xml.Name `xml:"body"`
	Section string   `xml:",innerxml"`
}

type XhtmlHead struct {
	XMLName xml.Name   `xml:"head"`
	Title   string     `xml:"title"`
	Meta    []MetaNode `xml:"meta"`
	Link    []Link     `xml:"link"`
}

type XhtmlHTML struct {
	XMLName xml.Name  `xml:"html"`
	Xmlns   string    `xml:"xmlns,attr"`
	Lang    string    `xml:"xml:lang,attr"`
	Header  XhtmlHead `xml:"head"`
	Body    XhtmlBody `xml:"body"`
}