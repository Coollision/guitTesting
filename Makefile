goApps := ${GOPATH}/bin/

all: local android


android:
	${goApps}fyne package -os android -appID com.example.myapp -icon Icon.png

runAndroid: android
	adb install ./guiTesting.apk

local:
	go build

run: local
	./guiTesting

ciTest:
	circleci local execute

installFyne:
	sudo apt-get install gcc libgl1-mesa-dev xorg-dev
	go get fyne.io/fyne/cmd/fyne@develop
