import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { GamePreviewComponent } from './game-preview/game-preview.component';
import { GameListComponent } from './game-list/game-list.component';

@NgModule({
    declarations: [
        AppComponent,
        GamePreviewComponent,
        GameListComponent
    ],
    imports: [
        BrowserModule
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule { }
