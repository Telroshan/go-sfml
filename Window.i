 %module "window"
 %{
 /* Includes the header in the wrapper code */
 #include <SFML/Window.h>
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

%include "Event.h"
%include "Event.h"
%include "Joystick.h"
%include "JoystickIdentification.h"
%include "Keyboard.h"
%include "Mouse.h"
%include "Sensor.h"
%include "Touch.h"
%include "Types.h"
%include "VideoMode.h"
%include "Window.h"
%include "WindowHandle.h"
