<template>
    <div class="w-48 flex flex-col">
        <ThemeToggle @toggledTheme="ThemeToggled" />
        <GameList />
        <Settings />
    </div>
    <div class="w-full h-screen flex flex-col self-center justify-center">
        <div class="h-16 bg-primary-light flex justify-center">
            <h1 class="text-xl self-center text-center">{{ GameTitle }}</h1>
        </div>
        <div class="flex-grow bg-primary-dark">
            <iframe ref="gamePage" :src="GetProject" allowtransparency="true" class="w-full h-full" frameborder="0"> </iframe>
        </div>
        <div class="h-12 relative flex justify-center bg-primary-light">
            <div class="absolute left-0 self-center px-4">
                <label class="mr-2"> Version: </label>

                <select class="cursor-pointer bg-primary hover:bg-primary-dark h-8 px-2 outline-none">
                    <option v-for="version in GameVersions" :key="version">{{ version }}</option>
                </select>
            </div>
            <button class="bg-tertiary hover:bg-primary-dark border-2 border-solid transform -translate-y-4 c-green">
                <h2 class="text-2xl font-bold px-12">Download</h2>
            </button>
        </div>
    </div>
</template>

<script setup>
import { ref, computed } from "@vue/reactivity";
import ThemeToggle from "./components/ThemeToggle.vue";
import GameList from "./components/GameList.vue";
import Settings from "./components/Settings.vue";
import { onMounted } from "@vue/runtime-core";

// CurrentGame
const GameTitle = ref("HyperHop");
const GameVersions = ref(["0.2.0 (Latest)", "0.1.0"]);

const gamePage = ref();

const isLight = ref(false);

const ThemeToggled = () => {
    isLight.value = localStorage.getItem("isLight");
};

onMounted(() => {
    ThemeToggled();
    var innerDoc = gamePage.value.contentDocument || gamePage.value.contentWindow.document;
    console.log(innerDoc.getElementById("app"));

    // eslint-disable-next-line no-unused-vars
    window.onerror = function (msg, _url, _line) {
        if (msg == "[IFRAME ERROR MESSAGE]") {
            return true;
        } else {
            //do nothing
        }
    };

    // disable right click to hide the fact its basicly just a browser window
    document.addEventListener("contextmenu", (event) => event.preventDefault());
    gamePage.value.contentWindow.document.addEventListener("contextmenu", (event) => event.preventDefault());
});

const GetProject = computed(() => {
    return "https://www.ethanconneely.com/projects/HyperHop?launcher=true&theme=" + (isLight.value == "true" ? "light" : "dark");
});
</script>
