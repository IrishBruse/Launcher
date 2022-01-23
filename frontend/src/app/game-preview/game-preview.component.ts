import { Component, OnInit, ChangeDetectionStrategy, Input, OnChanges, SimpleChanges } from '@angular/core';

@Component({
    selector: 'game-preview',
    templateUrl: './game-preview.component.html',
    styleUrls: ['./game-preview.component.css'],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class GamePreviewComponent implements OnInit, OnChanges {

    @Input()
    project: string = "";
    url: string = "";

    constructor() { }

    ngOnChanges(): void {
        this.updateIframe();
    }

    ngOnInit(): void {
        this.updateIframe();
    }

    updateIframe() {
        // This is unsafe but DomSanitizer wasnt working and this is a local app so it shouldnt matter
        const iframe = document.getElementById('game-preview-iframe') as HTMLIFrameElement;
        if (iframe.contentWindow != null) {
            iframe.contentWindow.location.replace('https://www.ethanconneely.com/projects/' + this.project + '/?launcher=true');
        }
    }

}
