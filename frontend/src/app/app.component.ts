import { ListItem } from 'src/app/App';
import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
    apps!: ListItem[];
    buttonText: string = "Download";
    currentProject!: ListItem;
    currentVersion: string = "";
    percent: number = 0.0;

    constructor() { }

    ngOnInit(): void {
        window.go.main.Launcher.GetApps().then(
            (apps) => {
                this.apps = JSON.parse(apps);
                this.currentProject = this.apps[0];
            }
        );


        window.runtime.EventsOn("downloadProgress",
            (percent) => {
                console.log("test");

                let btn = document.querySelector(".game-button");

                window.runtime.LogDebug(percent);

                let percentageText = btn?.querySelector(".app-button-text") as HTMLElement;
                percentageText.textContent = percent + "%";

                let fill = btn?.querySelector(".download-fill") as HTMLElement;
                fill.style.width = percent + "%";
            });


        this.updateIframe();
    }

    updateIframe() {
        // This is unsafe but DomSanitizer wasnt working and this is a local app so it shouldnt matter
        const iframe = document.getElementById('game-preview-iframe') as HTMLIFrameElement;
        if (iframe.contentWindow != null) {
            iframe.contentWindow.location.replace('https://www.ethanconneely.com/projects/' + this.currentProject + '/?launcher=true');
        }
    }

    getVersions(): string[] {
        if (this.apps !== undefined) {
            let t = this.apps.find((a) => a.Name == this.currentProject.Name)?.Versions;
            if (t === undefined) {
                return ["ERROR"];
            }
            return t;
        }
        return ["ERROR"];
    }

    modalClose() {
        let popup = document.getElementById("popup");
        if (popup != undefined) {
            popup.classList.add("modal-closed");
        }
    }

    modalOpen() {
        let popup = document.getElementById("popup");
        if (popup != undefined) {
            popup.classList.remove("modal-closed");
        }
    }

    modalDelete() {
        this.modalClose();
    }

    appInteract() {
        let btn = document.querySelector(".game-button")
        btn?.classList.toggle("downloading")

        console.log("interact");


        window.go.main.Launcher.Download(this.currentProject.Name + "/" + this.currentVersion);
    }

    selectApp(app: ListItem) {
        if (this.currentProject != app) {
            this.currentProject = app;
            this.updateIframe();
        }
    }

    selectVersion(selectElement: any) {
        console.log("test", selectElement);
    }
}
