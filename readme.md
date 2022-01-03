##Installation is as simple as running the following command (Go 1.17+):

```shell
go install github.com/wailsapp/wails/cmd/wails@latest
```

##Build it!
Change into the project directory cd my-project and compile your application using the build command wails build.

If all went well, you should have a compiled program in the build directory. Run it with build/my-project or double click myproject.exe if on windows.

##Serve
```shell
wails serve
```

While developing your apps using wails the preferred method is by the serve command wails serve.

This produces a much faster lightweight build in debug mode, excluding npm build scripts, saving time when developing the backend and also enabling use of npm run serve for partial browser development of frontend!

```shell
npm run serve
```