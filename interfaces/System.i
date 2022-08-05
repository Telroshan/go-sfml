 %module "system"
 %{
 /* Includes the header in the wrapper code */
 #include <SFML/System.h>
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

%include "Vector2.h"
%include "Vector3.h"
