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
    currentProject: string = "Monoboy";
    url: string = "";

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
        // if (this.apps !== undefined) {
        //     let t = this.apps.find((a) => a.name == this.currentProject)?.versions;
        //     if (t === undefined) {
        //         return ["ERROR"];
        //     }
        //     return t;
        // }
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
    }

    // GameListComponent
    selectApp(event: Event, app: ListItem) {
        let node = (event.target as HTMLElement);

        if (node.nodeName !== "BUTTON") {
            node = node.parentElement as HTMLElement;
        }

        document.querySelectorAll(".app-button").forEach((e) => {
            e.classList.remove("current")
        });

        let update = false;
        if (this.currentProject != app.Name) {
            update = true;
        }

        this.currentProject = app.Name;

        if (update) {
            this.updateIframe();
        }

        node.classList.toggle("current", true);

    }
}
