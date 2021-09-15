; -- Example2.iss --
; Same as Example1.iss, but creates its icon in the Programs folder of the
; Start Menu instead of in a subfolder, and also creates a desktop icon.

; SEE THE DOCUMENTATION FOR DETAILS ON CREATING .ISS SCRIPT FILES!

[Setup]
AppName=IrishBruse Launcher
AppVersion=1.0
WizardStyle=modern
SetupIconFile=Icon.ico
DefaultDirName={autopf}\IrishBruse Launcher
OutputBaseFilename="Launcher Installer"
; Since no icons will be created in "{group}", we don't need the wizard
; to ask for a Start Menu folder name:
DisableProgramGroupPage=yes
UninstallDisplayIcon={app}\Launcher.exe
Compression=lzma2

SolidCompression=yes
OutputDir=./Build/

[Files]
Source: "Launcher.exe"; DestDir: "{app}"

[Icons]
Name: "{autoprograms}\IrishBruse Launcher"; Filename: "{app}\Launcher.exe"
