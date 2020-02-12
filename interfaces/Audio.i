 %module "audio"
 %{
 /* Includes the header in the wrapper code */
#include <SFML/Audio.h>
 %}


%define linux %enddef

%include "SFML/Config.h"

%include "SFML/Audio/Export.h"
%include "SFML/Audio/Listener.h"
%include "SFML/Audio/Music.h"
%include "SFML/Audio/SoundBuffer.h"
%include "SFML/Audio/SoundBufferRecorder.h"
%include "SFML/Audio/Sound.h"
%include "SFML/Audio/SoundRecorder.h"
%include "SFML/Audio/SoundStatus.h"
%include "SFML/Audio/SoundStream.h"
%include "SFML/Audio/Types.h"

