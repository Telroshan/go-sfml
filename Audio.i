 %module "window"
 %{
 /* Includes the header in the wrapper code */
#include "../Audio.h"
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

%include "Listener.h"
%include "Music.h"
%include "Sound.h"
%include "SoundBuffer.h"
%include "SoundBufferRecorder.h"
%include "SoundRecorder.h"
%include "SoundStatus.h"
%include "SoundStream.h"
%include "Types.h"
