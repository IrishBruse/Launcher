<template>
    <div class="flex flex-col w-screen h-screen">
        <Titlebar :CurrentGame="currentGame"></Titlebar>

        <div class="w-full h-full flex-grow flex flex-row">
            <GameList @projectChanged="changeProject" />
            <ProjectDescription :CurrentGame="currentGame"></ProjectDescription>
        </div>

        <ProjectInteraction :CurrentGame="currentGame" @selectedVersionChanged="selectedVersionChanged" @deletePopup="deletePopup"></ProjectInteraction>
    </div>

    <div class="absolute w-screen h-screen bg-primary transition duration-200" ref="LoadingScreen">
        <div class="flex justify-center h-full">
            <h1 class="justify-start self-center px-2 text-4xl">Loading</h1>
            <div class="w-16 h-16 self-center">
                <h2 class="absolute icons text-6xl text-tertiary clipLoadingCircle animate-spin">circle</h2>
            </div>
        </div>
    </div>

    <div class="absolute shadow-md px-12 py-2 rounded-md bg-primary-light left-1/2 top-1/2 transform -translate-x-1/2 -translate-y-1/2 Hidden" id="modal">
        <h1 class="mx-auto my-4">Confirm Deletion?</h1>
        <div class="flex justify-between">
            <button class="rounded-md p-1 font-semibold text-primary bg-delete hover:bg-delete-hover" @click="Confirm()">Delete</button>
            <button class="rounded-md p-1 font-semibold text-primary bg-secondary hover:bg-secondary-dark" @click="Cancel()">Cancel</button>
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from "@vue/reactivity";
import ProjectDescription from "./components/ProjectView.vue";
import ProjectInteraction from "./components/ProjectInteraction.vue";
import Titlebar from "./components/Titlebar.vue";
import GameList from "./components/GameList.vue";
import { onMounted, watch } from "vue";

// CurrentGame
const finishedLoading = ref(false);
const currentGame = ref({ Name: null, Versions: [], Url: "" });
const selectedVersion = ref();
const LoadingScreen = ref();

var deletionConfirmedCallback;

const selectedVersionChanged = (version) => {
    selectedVersion.value = version;
};

const deletePopup = (deletionConfirmed) => {
    document.getElementById("modal").classList.remove("Hidden");

    deletionConfirmedCallback = deletionConfirmed;
};

const changeProject = (projectName, projectVersions) => {
    currentGame.value.Name = projectName;
    currentGame.value.Versions = projectVersions.reverse();
    currentGame.value.Url = "https://www.ethanconneely.com/projects/" + projectName + "/?launcher=true&theme=dark";
};

const Cancel = () => {
    document.getElementById("modal").classList.toggle("Hidden");
};

const Confirm = () => {
    document.getElementById("modal").classList.toggle("Hidden");
    removeItemOnce(DownloadedVersions[currentGame.value.Name], selectedVersion.value);
    deletionConfirmedCallback();
    Delete(currentGame.value.Name, selectedVersion.value);
};

function removeItemOnce(arr, value) {
    var index = arr.indexOf(value);
    if (index > -1) {
        arr.splice(index, 1);
    }
    return arr;
}

onMounted(() => {
    setTimeout(() => {
        start();
    }, 1);

    // disable right click to hide the fact its basicly just a browser window
    document.addEventListener("contextmenu", (event) => event.preventDefault());

    window.addEventListener("DownloadMetadataEvent", () => {
        finishedLoading.value = true;
        setTimeout(() => {
            LoadingScreen.value.classList.add("Hidden");
            setTimeout(() => {
                LoadingScreen.value.classList.add("Disabled");
                LoadingScreen.value.classList.remove("Hidden");
            }, 300);
        }, 3000);
    });
});
</script>

<style>
.Hidden {
    @apply opacity-0;
}

.Disabled {
    @apply hidden;
}
</style>
