# Tennis example

Go version of the [eponymous SFML example](https://github.com/SFML/SFML/tree/master/examples/tennis)

This code follows the same structure, call orders, with the same comments _(with just fixed maths lol)_ as the original, so it should give you a good idea on how to transpose C++ SFML code to go-sfml!

1. Open `build.bat` script and update its `CSFML_PATH` variable to match the path of your CSFML installation 
2. Run the `build.bat` script
3. Please note that you may need the `openal32.dll` to run the executable, you'll find it in SFML's _(not CSFML)_ `bin` folder
4. Run `tennis.exe`, the following result should appear!

![tennis](https://user-images.githubusercontent.com/19146183/182578103-eaef7229-0cc2-43a1-b184-9bf4950de8e7.gif)

If you're having issues building & running this example, please read the [general README](https://github.com/Telroshan/go-sfml#usage), and feel free to open an issue if your problem is not listed in here.