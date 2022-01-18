import { Component, OnInit, ChangeDetectionStrategy, Input } from '@angular/core';
import { DomSanitizer, SafeResourceUrl, } from '@angular/platform-browser';

@Component({
    selector: 'game-preview',
    templateUrl: './game-preview.component.html',
    styleUrls: ['./game-preview.component.css'],
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class GamePreviewComponent implements OnInit {

    @Input()
    project!: string;

    url!: SafeResourceUrl;

    constructor(public sanitizer: DomSanitizer) { }

    ngOnInit(): void {
        this.url = this.sanitizer.bypassSecurityTrustResourceUrl('https://www.ethanconneely.com/projects/' + this.project + '/?launcher=true')
    }

}
