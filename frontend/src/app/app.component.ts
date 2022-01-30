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
    currentProject: string = "HyperHop";
    currentVersion: string = "";
    percent: number = 10.0;

    constructor() { }

    ngOnInit(): void {
        window.go.main.Launcher.GetApps().then(
            (apps) => {
                this.apps = JSON.parse(apps);
                console.log(this.apps);
            }
        )

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
            let t = this.apps.find((a) => a.Name == this.currentProject)?.Versions;
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

        let percentageText = btn?.querySelector(".app-button-text") as HTMLElement;
        percentageText.textContent = percent + "%";

        let fill = btn?.querySelector(".download-fill") as HTMLElement;
        fill.style.width = percent + "%";
    }

    selectApp(app: ListItem) {
        if (this.currentProject != app.Name) {
            this.currentProject = app.Name;
            this.updateIframe();
        }
    }

    selectVersion(selectElement: any) {
        console.log("test", selectElement);
    }
}
