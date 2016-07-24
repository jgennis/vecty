//go:generate go run generate.go

// Package svg defines markup to create DOM elements.
//
// Generated from "SVG element reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element, licensed under
// CC-BY-SA 2.5.
package svg

import "github.com/gopherjs/vecty"

// The SVG Anchor Element () defines a hyperlink
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/a
func Anchor(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "a", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The altGlyph element allows sophisticated selection of the glyphs used to
// render its child character data.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/altGlyph
func AlternateGlyph(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "altGlyph", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The altGlyphDef element defines a substitution representation for glyphs.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/altGlyphDef
func AlternateGlyphDefinition(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "altGlyphDef", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The altGlyphItem element provides a set of candidates for glyph substitution
// by the <altGlyph> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/altGlyphItem
func AlternateGlyphItem(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "altGlyphItem", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The animate element is put inside a shape element and defines how an
// attribute of an element changes over the animation. The attribute will
// change from the initial value to the end value in the duration specified.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animate
func Animate(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "animate", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The animateColor element specifies a color transformation over time.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateColor
func AnimateColor(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "animateColor", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The animateMotion element causes a referenced element to move along a motion
// path.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateMotion
func AnimateMotion(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "animateMotion", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The animateTransform element animates a transformation attribute on a target
// element, thereby allowing animations to control translation, scaling,
// rotation and/or skewing.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/animateTransform
func AnimateTransform(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "animateTransform", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The circle element is an SVG basic shape, used to create circles based on a
// center point and a radius.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/circle
func Circle(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "circle", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The clipping path restricts the region to which paint can be applied.
// Conceptually, any parts of the drawing that lie outside of the region
// bounded by the currently active clipping path are not drawn.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/clipPath
func ClipPath(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "clipPath", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The element allows describing the color profile used for the image.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/color-profile
func ColorProfile(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "color-profile", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The cursor element can be used to define a platform-independent custom
// cursor. A recommended approach for defining a platform-independent custom
// cursor is to create a PNG image and define a cursor element that references
// the PNG image and identifies the exact position within the image which is
// the pointer position (i.e., the hot spot).
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/cursor
func Cursor(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "cursor", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// SVG allows graphical objects to be defined for later reuse. It is
// recommended that, wherever possible, referenced elements be defined inside
// of a defs element. Defining these elements inside of a defs element promotes
// understandability of the SVG content and thus promotes accessibility.
// Graphical elements defined in a defs will not be directly rendered. You can
// use a <use> element to render those elements wherever you want on the
// viewport.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/defs
func Definitions(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "defs", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// Each container element or graphics element in an SVG drawing can supply a
// desc description string where the description is text-only. When the current
// SVG document fragment is rendered as SVG on visual media, desc elements are
// not rendered as part of the graphics. Alternate presentations are possible,
// both visual and aural, which display the desc element but do not display
// path elements or other graphics elements. The desc element generally improve
// accessibility of SVG documents
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/desc
func Description(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "desc", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The ellipse element is an SVG basic shape, used to create ellipses based on
// a center coordinate, and both their x and y radius.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/ellipse
func Ellipse(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "ellipse", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The feBlend filter composes two objects together ruled by a certain blending
// mode. This is similar to what is known from image editing software when
// blending two layers. The mode is defined by the mode attribute.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feBlend
func FilterEffectBlend(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feBlend", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter changes colors based on a transformation matrix. Every pixel's
// color value (represented by an [R,G,B,A] vector) is matrix multiplied to
// create a new color.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feColorMatrix
func FilterEffectColorMatrix(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feColorMatrix", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The color of each pixel is modified by changing each channel (R, G, B, and
// A) to the result of what the children <feFuncR>, <feFuncB>, <feFuncG>, and
// <feFuncA> return.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComponentTransfer
func FilterEffectComponentTransfer(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feComponentTransfer", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive performs the combination of two input images
// pixel-wise in image space using one of the Porter-Duff compositing
// operations: over, in, atop, out, xor. Additionally, a component-wise
// arithmetic operation (with the result clamped between [0..1]) can be
// applied.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feComposite
func FilterEffectComposite(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feComposite", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// the feConvolveMatrix element applies a matrix convolution filter effect. A
// convolution combines pixels in the input image with neighboring pixels to
// produce a resulting image. A wide variety of imaging operations can be
// achieved through convolutions, including blurring, edge detection,
// sharpening, embossing and beveling.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feConvolveMatrix
func FeConvolveMatrix(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feConvolveMatrix", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive lights an image using the alpha channel as a bump map.
// The resulting image, which is an RGBA opaque image, depends on the light
// color, light position and surface geometry of the input bump map.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDiffuseLighting
func FeDiffuseLighting(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feDiffuseLighting", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive uses the pixels values from the image from in2 to
// spatially displace the image from in.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDisplacementMap
func FilterEffectDisplacementMap(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feDisplacementMap", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive define a distant light source that can be used within
// a lighting filter primitive : <fediffuselighting> or <fespecularlighting>.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feDistantLight
func FilterEffectDistantLight(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feDistantLight", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The filter fills the filter subregion with the color and opacity defined by
// flood-color and flood-opacity.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFlood
func FilterEffectFlood(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feFlood", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive defines the transfer function for the alpha component
// of the input graphic of its parent <fecomponenttransfer> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncA
func FilterEffectFunctionA(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feFuncA", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive defines the transfer function for the blue component
// of the input graphic of its parent <fecomponenttransfer> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncB
func FilterEffectFunctionB(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feFuncB", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive defines the transfer function for the green component
// of the input graphic of its parent <fecomponenttransfer> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncG
func FeFuncG(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feFuncG", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive defines the transfer function for the red component of
// the input graphic of its parent <fecomponenttransfer> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feFuncR
func FeFuncR(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feFuncR", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The filter blurs the input image by the amount specified in stdDeviation,
// which defines the bell-curve.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feGaussianBlur
func FilterEffectGaussianBlur(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feGaussianBlur", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The feImage filter fetches image data from an external source and provides
// the pixel data as output (meaning if the external source is an SVG image, it
// is rasterized.)
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feImage
func FilterEffectImage(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feImage", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The feMerge filter allows filter effects to be applied concurrently instead
// of sequentially. This is achieved by other filters storing their output via
// the result attribute and then accessing it in a <femergenode> child.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMerge
func FilterEffectMerge(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feMerge", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The feMergeNode takes the result of another filter to be processed by its
// parent <femerge>.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMergeNode
func FilterEffectMergeNode(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feMergeNode", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter is used to erode or dilate the input image. It's usefulness lies
// especially in fattening or thinning effects.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feMorphology
func FilterEffectMorphology(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feMorphology", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The input image as a whole is offset by the values specified in the dx and
// dy attributes. It's used in creating drop-shadows.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feOffset
func FilterEffectOffset(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feOffset", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This element implements the SVGFEPointLightElement interface.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/fePointLight
func FilterEffectPointLight(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "fePointLight", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive lights a source graphic using the alpha channel as a
// bump map. The resulting image is an RGBA image based on the light color. The
// lighting calculation follows the standard specular component of the Phong
// lighting model. The resulting image depends on the light color, light
// position and surface geometry of the input bump map. The result of the
// lighting calculation is added. The filter primitive assumes that the viewer
// is at infinity in the z direction.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpecularLighting
func FilterEffectSpecularLighting(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feSpecularLighting", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The feSpotLight element is one of the ligth source elements used for SVG
// files.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feSpotLight
func FeSpotLight(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feSpotLight", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// An input image is tiled and the result used to fill a target. The effect is
// similar to the one of a <pattern>.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feTile
func FilterEffectTile(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feTile", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// This filter primitive creates an image using the Perlin turbulence function.
// It allows the synthesis of artificial textures like clouds or marble.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/feTurbulence
func FilterEffectTurbulence(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "feTurbulence", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The filter element serves as container for atomic filter operations. It is
// never rendered directly. A filter is referenced by using the filter
// attribute on the target SVG element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/filter
func Filter(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "filter", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The font element defines a font to be used for text layout.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font
func Font(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "font", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The font-face element corresponds to the CSS @font-face declaration. It
// defines a font's outer properties.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face
func FontFace(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "font-face", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The font-face-format element describes the type of font referenced by its
// parent <font-face-uri>.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-format
func FontFaceFormat(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "font-face-format", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The font-face-name element points to a locally installed copy of this font,
// identified by its name.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-name
func FontFaceName(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "font-face-name", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The font-face-src element corresponds to the src property in CSS @font-face
// descriptions. It serves as container for <font-face-name>, pointing to
// locally installed copies of this font, and <font-face-uri>, utilizing
// remotely defined fonts.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-src
func FontFaceSource(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "font-face-src", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The font-face-uri element points to a remote definition of the current font.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/font-face-uri
func FontFaceURI(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "font-face-uri", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The foreignObject element allows for inclusion of a foreign XML namespace
// which has its graphical content drawn by a different user agent. The
// included foreign graphical content is subject to SVG transformations and
// compositing.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/foreignObject
func ForeignObject(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "foreignObject", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The g element is a container used to group objects. Transformations applied
// to the g element are performed on all of its child elements. Attributes
// applied are inherited by child elements. In addition, it can be used to
// define complex objects that can later be referenced with the <use> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/g
func Group(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "g", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// A glyph defines a single glyph in an SVG font.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/glyph
func Glyph(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "glyph", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The glyphRef element provides a single possible glyph to the referencing
// <altGlyph> substitution.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/glyphRef
func GlyphReference(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "glyphRef", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The horizontal distance between two glyphs can be fine-tweaked with an hkern
// Element. This process is known as Kerning.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/hkern
func HorizontalKern(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "hkern", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The SVG Image Element () allows a raster image into be included in an SVG
// document.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/image
func Image(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "image", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The line element is an SVG basic shape, used to create a line connecting two
// points.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/line
func Line(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "line", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The linearGradient element lets authors define linear gradients to fill or
// stroke graphical elements.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/linearGradient
func LinearGradient(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "linearGradient", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The marker element defines the graphics that is to be used for drawing
// arrowheads or polymarkers on a given <path>, <line>, <polyline> or <polygon>
// element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/marker
func Marker(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "marker", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// In SVG, you can specify that any other graphics object or <g> element can be
// used as an alpha mask for compositing the current object into the
// background. A mask is defined with the mask element. A mask is
// used/referenced using the mask property.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mask
func Mask(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "mask", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// Metadata is structured data about data. Metadata which is included with SVG
// content should be specified within metadata elements. The contents of the
// metadata should be elements from other XML namespaces such as RDF, FOAF,
// etc.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/metadata
func Metadata(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "metadata", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The missing-glyph's content is rendered, if for a given character the font
// doesn't define an appropriate <glyph>.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/missing-glyph
func MissingGlyph(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "missing-glyph", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// the mpath sub-element for the <animatemotion> element provides the ability
// to reference an external <path> element as the definition of a motion path.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/mpath
func MotionPath(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "mpath", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The path element is the generic element to define a shape. All the basic
// shapes can be created with a path element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/path
func Path(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "path", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// A pattern is used to fill or stroke an object using a pre-defined graphic
// object which can be replicated ("tiled") at fixed intervals in x and y to
// cover the areas to be painted. Patterns are defined using the pattern
// element and then referenced by properties fill and stroke on a given
// graphics element to indicate that the given element shall be filled or
// stroked with the referenced pattern.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/pattern
func Pattern(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "pattern", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The polygon element defines a closed shape consisting of a set of connected
// straight line segments. The last point is connected to the first point. For
// open shapes see the <polyline> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/polygon
func Polygon(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "polygon", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The polyline element is an SVG basic shape, used to create a series of
// straight lines connecting several points. Typically a polyline is used to
// create open shapes as the last point is not connected to the first point.
// For closed shapes see the <polygon> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/polyline
func Polyline(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "polyline", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// radialGradient lets authors define radial gradients to fill or stroke
// graphical elements.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/radialGradient
func RadialGradient(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "radialGradient", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The rect element is an SVG basic shape, used to create rectangles based on
// the position of a corner and their width and height. It may also be used to
// create rectangles with rounded corners.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/rect
func Rectangle(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "rect", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// A SVG script element is equivalent to the script element in HTML and thus is
// the place for scripts (e.g., ECMAScript).
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/script
func Script(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "script", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The set element provides a simple means of just setting the value of an
// attribute for a specified duration. It supports all attribute types,
// including those that cannot reasonably be interpolated, such as string and
// boolean values. The set element is non-additive. The additive and accumulate
// attributes are not allowed, and will be ignored if specified.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/set
func Set(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "set", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The ramp of colors to use on a gradient is defined by the stop elements that
// are child elements to either the <linearGradient> element or the
// <radialGradient> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/stop
func Stop(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "stop", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The style element allows style sheets to be embedded directly within SVG
// content. SVG's style element has the same attributes as the corresponding
// element in HTML (see HTML's <style> element).
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/style
func Style(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "style", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The svg element can be used to nest a standalone SVG fragment inside the
// current document (for example an HTML document) as long as the svg is not
// the root element. This standalone fragment has its own viewport and
// coordinate system.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/svg
func SVG(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "svg", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The switch element evaluates the requiredFeatures, requiredExtensions and
// systemLanguage attributes on its direct child elements in order, and then
// processes and renders the first child for which these attributes evaluate to
// true. All others will be bypassed and therefore not rendered. If the child
// element is a container element such as a <g>, then the entire subtree is
// either processed/rendered or bypassed/not rendered.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/switch
func Switch(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "switch", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The symbol element is used to define graphical template objects which can be
// instantiated by a <use> element. The use of symbol elements for graphics
// that are used multiple times in the same document adds structure and
// semantics. Documents that are rich in structure may be rendered graphically,
// as speech, or as braille, and thus promote accessibility. note that a symbol
// element itself is not rendered. Only instances of a symbol element (i.e., a
// reference to a symbol by a <use> element) are rendered.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/symbol
func Symbol(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "symbol", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The text element defines a graphics element consisting of text. Note that it
// is possible to apply a gradient, pattern, clipping path, mask or filter to
// text
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/text
func Text(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "text", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// In addition to text drawn in a straight line, SVG also includes the ability
// to place text along the shape of a <path> element. To specify that a block
// of text is to be rendered along the shape of a <path>, include the given
// text within a textPath element which includes an xlink:href attribute with a
// reference to a <path> element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/textPath
func TextPath(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "textPath", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// Each container element or graphics element in an SVG drawing can supply a
// title description string where the description is text-only. When the
// current SVG document fragment is rendered as SVG on visual media, title
// element is not rendered as part of the graphics. However, some user agents
// may, for example, display the title element as a tooltip. Alternate
// presentations are possible, both visual and aural, which display the title
// element but do not display path elements or other graphics elements. The
// title element generally improve accessibility of SVG documents
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/title
func Title(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "title", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The textual content for a <text> can be either character data directly
// embedded within the <text> element or the character data content of a
// referenced element, where the referencing is specified with a tref element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/tref
func TextReference(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "tref", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// Within a <text> element, text and font properties and the current text
// position can be adjusted with absolute or relative coordinate values by
// including a tspan element.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/tspan
func TextSpan(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "tspan", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The use element takes nodes from within the SVG document, and duplicates
// them somewhere else. The effect is the same as if the nodes were deeply
// cloned into a non-exposed DOM, and then pasted where the use element is,
// much like cloned template elements in HTML5. Since the cloned nodes are not
// exposed, care must be taken when using CSS to style a use element and its
// hidden descendants. CSS attributes are not guaranteed to be inherited by the
// hidden, cloned DOM unless you explicitly request it using CSS inheritance.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/use
func Use(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "use", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// A view is a defined way to view the image, like a zoom level or a detail
// view.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/view
func View(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "view", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}

// The vertical distance between two glyphs in top-to-bottom fonts can be
// fine-tweaked with an vkern Element. This process is known as Kerning.
//
// https://developer.mozilla.org/en-US/docs/Web/SVG/Element/vkern
func VerticalKern(markup ...vecty.Markup) *vecty.Element {
	e := &vecty.Element{TagName: "vkern", Namespace: "http://www.w3.org/2000/svg"}
	vecty.List(markup).Apply(e)
	return e
}
