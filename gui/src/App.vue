<template>
    <div class="flex flex-col w-screen h-screen">
        <Titlebar :CurrentGame="currentGame"></Titlebar>

        <div class="w-full h-full flex-grow flex flex-row">
            <GameList @projectChanged="changeProject" />
            <ProjectDescription :CurrentGame="currentGame"></ProjectDescription>
        </div>

        <ProjectInteraction :CurrentGame="currentGame"></ProjectInteraction>
    </div>

    <div class="absolute w-screen h-screen bg-primary transition duration-200" ref="LoadingScreen">
        <div class="flex justify-center h-full">
            <h1 class="justify-start self-center px-2 text-4xl">Loading</h1>
            <div class="w-16 h-16 self-center">
                <h2 class="absolute icons text-6xl text-tertiary clipLoadingCircle animate-spin">circle</h2>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from "@vue/reactivity";
import ProjectDescription from "./components/MainColumn/ProjectView.vue";
import ProjectInteraction from "./components/MainColumn/ProjectInteraction.vue";
import Titlebar from "./components/MainColumn/Titlebar.vue";
import GameList from "./components/Sidebar/GameList.vue";
import { onMounted, watch } from "vue";

// CurrentGame
const finishedLoading = ref(false);
const currentGame = ref({ Name: null, Versions: [], Url: "" });
const LoadingScreen = ref();

const changeProject = (projectName, projectVersions) => {
    currentGame.value.Name = projectName;
    currentGame.value.Versions = projectVersions.reverse();
    currentGame.value.Url = "https://www.ethanconneely.com/projects/" + projectName + "/?launcher=true&theme=dark";
};

onMounted(() => {
    setTimeout(() => {
        start();
    }, 1);

    // disable right click to hide the fact its basicly just a browser window
    document.addEventListener("contextmenu", (event) => event.preventDefault());

    window.addEventListener("DownloadMetadataEvent", () => {
        finishedLoading.value = true;
        setTimeout(() => {
            LoadingScreen.value.classList.add("Hide");
            setTimeout(() => {
                LoadingScreen.value.classList.add("Disabled");
                LoadingScreen.value.classList.remove("Hide");
            }, 300);
        }, 3000);
    });
});
</script>

<style>
.Hide {
    @apply opacity-0;
}

.Disabled {
    @apply hidden;
}
</style>
