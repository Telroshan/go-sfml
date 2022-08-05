 %module "graphics"
 %{
 /* Includes the header in the wrapper code */
 #include <SFML/System.h>
 #include <SFML/Graphics.h>
 %}

 /* Parse the header file to generate wrappers */

%define MACOSX %enddef
%define __APPLE__ %enddef
%define MACOSX %enddef
%define macintosh %enddef
%define Macintosh %enddef
%define CSFML_SYSTEM_MACOS %enddef

%ignore sfRenderWindow_capture;
%ignore sfShader_setFloatParameter;
%ignore sfShader_setFloat2Parameter;
%ignore sfShader_setFloat3Parameter;
%ignore sfShader_setFloat4Parameter;
%ignore sfShader_setVector2Parameter;
%ignore sfShader_setVector3Parameter;
%ignore sfShader_setColorParameter;
%ignore sfShader_setTransformParameter;
%ignore sfShader_setTextureParameter;
%ignore sfShader_setCurrentTextureParameter;

%include <SFML/Config.h>
%include "Export.h"

%include <SFML/System/Vector2.h>
%include <SFML/System/Vector3.h>

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
%include "VertexBuffer.h"
%include "View.h"
