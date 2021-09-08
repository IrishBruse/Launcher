<template>
    <div class="w-48 flex flex-col">
        <ThemeToggle @toggledTheme="toggledTheme" />
        <GameList @projectChanged="changeProject" />
        <Settings />
    </div>
    <div class="w-full h-screen flex flex-col self-center justify-center">
        <Titlebar :CurrentGame="currentGame"></Titlebar>
        <ProjectDescription :CurrentGame="currentGame"></ProjectDescription>
        <ProjectInteraction :CurrentGame="currentGame"></ProjectInteraction>
    </div>
</template>

<script setup>
import { ref, computed } from "@vue/reactivity";
import ProjectDescription from "./components/MainColumn/ProjectView.vue";
import ProjectInteraction from "./components/MainColumn/ProjectInteraction.vue";
import Titlebar from "./components/MainColumn/Titlebar.vue";
import ThemeToggle from "./components/Sidebar/ThemeToggle.vue";
import GameList from "./components/Sidebar/GameList.vue";
import Settings from "./components/Sidebar/Settings.vue";
import { onMounted } from "vue";

// CurrentGame
const finishedLoading = ref(false);
const currentGame = ref({ Name: null, Versions: [], Url: "" });

const isLight = ref(false);

const changeProject = (projectName, projectVersions) => {
    currentGame.value.Name = projectName;
    currentGame.value.Versions = projectVersions.reverse();
    currentGame.value.Url = "https://www.ethanconneely.com/projects/" + projectName + "/?launcher=true&theme=" + (isLight.value == "true" ? "light" : "dark");
};

const toggledTheme = () => {
    isLight.value = localStorage.getItem("isLight");
};

onMounted(() => {
    setTimeout(() => {
        start();
    }, 1);
    toggledTheme();

    // disable right click to hide the fact its basicly just a browser window
    window.addEventListener("contextmenu", (event) => event.preventDefault());

    window.addEventListener("DownloadedProjectsMetadata", () => {
        finishedLoading.value = true;
    });

    // TODO: When Downloading a game pop this up
    window.addEventListener("beforeunload", function (e) {
        // e.preventDefault();
        // e.returnValue = "";
    });
});
</script>
