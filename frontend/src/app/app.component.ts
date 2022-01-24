import { Component, OnInit } from '@angular/core';
import { App } from './App';

@Component({
    selector: 'app-root',
    templateUrl: './app.component.html',
    styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
    apps!: App[];
    buttonText: string = "Download";
    currentProject: string = "Monoboy";
    url: string = "";

    constructor() { }

    ngOnInit(): void {

        // window.go.main.App.GetAppIcons().then(
        //     (icons) => {
        //         icons.forEach(icon => {
        //             console.log(icon);
        //         });
        //     }
        // )

        this.apps = [
            {
                name: "test",
                iconUrl: "https://ucf2f66ac0d8f061d8d9cb8de8fe.dl.dropboxusercontent.com/cd/0/get/BeWsTrdaw-EWNBt5bEb9S8X-L4xHMcJvOKfaGRPMrfxag1cuIpCPEHVREQesh_o2LTnTF41JQXMAkN1jgF2JdcPqiTU9VsD3_aTLOVmgzGev5fSxMzOBxVq5w8XErS3jjWAM3X5KJHusyJQJWw8ImdGH/file",
                versions: ["0.1.0", "0.2.0", "0.3.0"]
            },
            {
                name: "Monoboy",
                iconUrl: "https://ucf2f66ac0d8f061d8d9cb8de8fe.dl.dropboxusercontent.com/cd/0/get/BeWsTrdaw-EWNBt5bEb9S8X-L4xHMcJvOKfaGRPMrfxag1cuIpCPEHVREQesh_o2LTnTF41JQXMAkN1jgF2JdcPqiTU9VsD3_aTLOVmgzGev5fSxMzOBxVq5w8XErS3jjWAM3X5KJHusyJQJWw8ImdGH/file",
                versions: ["0.1.0", "0.2.0", "0.3.0", "0.4.0"]
            },
            {
                name: "Top Down Shooter",
                iconUrl: "https://ucf2f66ac0d8f061d8d9cb8de8fe.dl.dropboxusercontent.com/cd/0/get/BeWsTrdaw-EWNBt5bEb9S8X-L4xHMcJvOKfaGRPMrfxag1cuIpCPEHVREQesh_o2LTnTF41JQXMAkN1jgF2JdcPqiTU9VsD3_aTLOVmgzGev5fSxMzOBxVq5w8XErS3jjWAM3X5KJHusyJQJWw8ImdGH/file",
                versions: ["0.1.0", "0.2.0"]
            }
        ]
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
        let t = this.apps.find((a) => a.name == this.currentProject)?.versions;
        if (t === undefined) {
            return ["ERROR"];
        }
        return t;
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
    selectApp(event: Event, app: App) {
        let node = (event.target as HTMLElement);

        if (node.nodeName !== "BUTTON") {
            node = node.parentElement as HTMLElement;
        }

        document.querySelectorAll(".app-button").forEach((e) => {
            e.classList.remove("current")
        });

        let update = false;
        if (this.currentProject != app.name) {
            update = true;
        }

        this.currentProject = app.name;

        if (update) {
            this.updateIframe();
        }

        node.classList.toggle("current", true);

    }
}
