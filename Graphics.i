 %module "graphics"
 %{
 /* Includes the header in the wrapper code */
#include "../Graphics.h"
 %}

 /* Parse the header file to generate wrappers */

%define MACOSX %enddef
%define __APPLE__ %enddef
%define MACOSX %enddef
%define macintosh %enddef
%define Macintosh %enddef
%define CSFML_SYSTEM_MACOS %enddef

%include <SFML/Config.h>

%include "Export.h"

%include "BlendMode.h"
%include "CircleShape.h"
%include "Color.h"
%include "ConvexShape.h"
%include "Font.h"
%include "FontInfo.h"
%include "Glsl.h"
%include "Glyph.h"
%include "Image.h"
%include "PrimitiveType.h"
%include "Rect.h"
%include "RectangleShape.h"
%include "RenderStates.h"
%include "RenderTexture.h"
%include "RenderWindow.h"
%include "Shader.h"
%include "Shape.h"
%include "Sprite.h"
%include "Text.h"
%include "Texture.h"
%include "Transform.h"
%include "Transformable.h"
%include "Types.h"
%include "Vertex.h"
%include "VertexArray.h"
%include "View.h"
