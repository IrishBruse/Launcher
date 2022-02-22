import { ListItem } from "./ListItem";
import { Component, ElementRef, OnInit, ViewChild } from "@angular/core";
import { AppState } from "./AppState";

@Component({ selector: "app-root", templateUrl: "./app.component.html", styleUrls: ["./app.component.css"] })
export class AppComponent implements OnInit
{
    apps = new Array<ListItem>();

    state = AppState.Download;

    currentVersion = "";
    currentProject = new ListItem();

    @ViewChild("buttonText") buttonText!: ElementRef<HTMLInputElement>;

    ngOnInit(): void
    {
        window.go.main.Launcher.GetApps().then(
            (apps) =>
            {
                this.apps = JSON.parse(apps);
                this.selectApp(this.apps[0]);
            }
        );

        window.runtime.EventsOn("downloadProgress",
            (percent: number) =>
            {
                this.buttonText.nativeElement.textContent = percent + "%";

                let btn = document.querySelector(".game-button");

                let fill = btn?.querySelector(".download-fill") as HTMLElement;
                fill.style.width = percent + "%";

                if (percent == 100)
                {
                    setTimeout(() =>
                    {
                        let btn = document.querySelector(".game-button");
                        if (btn != null)
                        {
                            btn.classList.remove("downloading");
                            this.changeState(AppState.Play);
                        }

                        this.currentProject.Downloaded.push(this.currentVersion);

                        fill.style.width = 0 + "%";
                    }, 1000);
                }

            }
        );
    }

    updateIframe()
    {
        // This is unsafe but DomSanitizer wasnt working and this is a local app so it shouldnt matter
        const iframe = document.getElementById("game-preview-iframe") as HTMLIFrameElement;
        if (iframe.contentWindow != null)
        {
            iframe.contentWindow.location.replace("https://www.ethanconneely.com/projects/" + this.currentProject.Name + "/?launcher=true");
        }
    }

    getVersions(): string[] | undefined
    {
        return this.apps.find((a) => a.Name == this.currentProject.Name)?.Versions;
    }

    modalToggle()
    {
        if (this.state == AppState.Playing || this.state == AppState.Downloading)
        {
            return;
        }

        let popup = document.getElementById("popup");
        if (popup != undefined)
        {
            popup.classList.toggle("modal-closed");
        }
    }

    modalDelete()
    {
        this.modalToggle();
        window.go.main.Launcher.Delete(this.currentProject.Name + "/" + this.currentVersion).then(() =>
        {
            this.currentProject.Downloaded.splice(this.currentProject.Downloaded.indexOf(this.currentVersion), 1);
        });
    }

    appInteract()
    {
        switch (this.state)
        {
        case AppState.Download:
            this.download();
            break;

        case AppState.Play:
            this.play();
            break;
        }
    }

    play()
    {
        this.changeState(AppState.Playing);
        window.go.main.Launcher.Play(this.currentProject.Name + "/" + this.currentVersion).then(() =>
        {
            this.onChangeVersion();
        });
    }

    download()
    {
        let btn = document.querySelector(".game-button");
        if (btn != null)
        {
            btn.classList.add("downloading");
            this.changeState(AppState.Downloading);
            this.buttonText.nativeElement.textContent = "0%";
        }

        window.go.main.Launcher.Download("/" + this.currentProject.Name + "/" + this.currentVersion + ".zip");
    }

    onChangeVersion()
    {
        window.runtime.LogInfo(this.currentVersion);

        this.changeState(AppState.Download);

        this.apps.forEach(app =>
        {
            if (app.Downloaded != null)
            {
                console.log(app);

                if (app.Downloaded.indexOf(this.currentVersion) !== -1)
                {
                    this.changeState(AppState.Play);
                }
            }
        });


        console.log("test");

    }

    private changeState(state: AppState)
    {
        this.state = state;

        let val = "Error";
        switch (state)
        {
        case AppState.Download:
            val = "Download";
            break;

        case AppState.Downloading:
            val = "Downloading";
            break;

        case AppState.Play:
            val = "Play";
            break;

        case AppState.Playing:
            val = "Playing";
            break;
        }
        this.buttonText.nativeElement.textContent = val;
    }

    selectApp(app: ListItem)
    {
        if (this.state == AppState.Playing || this.state == AppState.Downloading)
        {
            return;
        }

        if (this.currentProject != app)
        {
            this.currentProject = app;
            this.currentVersion = app.Versions[0];
            this.onChangeVersion();
            this.updateIframe();
        }
    }
}
