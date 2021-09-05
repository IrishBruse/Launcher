<template>
    <div class="w-48 flex flex-col">
        <ThemeToggle @toggledTheme="toggledTheme" />
        <GameList @projectChanged="changeProject" />
        <Settings />
    </div>
    <div class="w-full h-screen flex flex-col self-center justify-center">
        <div class="h-16 bg-primary-dark flex justify-center">
            <h1 class="text-xl self-center text-center">{{ GameTitle }}</h1>
        </div>
        <div class="flex-grow bg-primary">
            <iframe v-if="GameTitle != null && GameTitle != 'Pick a game'" ref="gamePage" :src="CurrentProjectUrl" class="w-full h-full" frameborder="0">
            </iframe>
            <div v-else-if="GameTitle != 'Pick a game'" class="flex justify-center h-full">
                <h2 class="justify-start self-center px-2">Loading</h2>
                <div class="w-8 h-8 self-center">
                    <h2 class="absolute icons text-2xl font-bold text-secondary-dark">radio_button_unchecked</h2>
                    <h2 class="absolute icons text-2xl clipLoadingCircle animate-spin">circle</h2>
                </div>
            </div>
        </div>
        <div class="h-12 relative flex justify-center bg-primary-dark">
            <div class="absolute left-0 self-center px-4">
                <label class="mr-2"> Version: </label>

                <select class="transition duration-300 cursor-pointer bg-primary hover:bg-primary-hover h-8 px-2 outline-none rounded-md">
                    <option v-for="(index, version) in GameVersions" :key="index">{{ version }}</option>
                </select>
            </div>
            <button class="transition duration-300 group bg-tertiary shadow-xl hover:bg-secondary hover:text-primary-light transform -translate-y-4 rounded-lg">
                <h2 class="text-2xl transition font-bold group-hover:text-primary-light px-12">Download</h2>
            </button>
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from "@vue/reactivity";
import ThemeToggle from "./components/ThemeToggle.vue";
import GameList from "./components/GameList.vue";
import Settings from "./components/Settings.vue";
import { onMounted } from "vue";

// CurrentGame
const GameTitle = ref(null);
const GameVersions = ref([]);

const gamePage = ref();

const isLight = ref(false);

const changeProject = (projectName, projectVersions) => {
    console.log(projectName);
    GameTitle.value = projectName;
    GameVersions.value = projectVersions;
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
    document.addEventListener("contextmenu", (event) => event.preventDefault());
});

const CurrentProjectUrl = computed(() => {
    return "https://www.ethanconneely.com/projects/" + GameTitle.value + "/?launcher=true&theme=" + (isLight.value == "true" ? "light" : "dark");
});
</script>
