<template>
    <a class="absolute w-10 h-10 self-center select-none cursor-pointer" @click="toggleDarkMode">
        <span class="icons hover:text-secondary absolute text-3xl transition duration-400" :class="darkMode ? 'opacity-0 z-0' : 'opacity-100 z-10'">
            dark_mode
        </span>
        <span class="icons hover:text-secondary absolute text-3xl transition duration-400" :class="darkMode ? 'opacity-100 z-10' : 'opacity-0 z-0'">
            light_mode
        </span>
    </a>
</template>

<script setup>
import { ref, onMounted } from "vue";

const darkMode = ref(true);

const toggleDarkMode = () => {
    darkMode.value = !darkMode.value;
    localStorage.setItem("isLight", darkMode.value);
    var body = document.getElementsByTagName("body")[0];
    body.classList.toggle("light", darkMode.value);
};

onMounted(() => {
    if (localStorage.getItem("isLight") === null) {
        var val = window.matchMedia("(prefers-color-scheme: dark)").matches;
        localStorage.setItem("isLight", val);
        darkMode.value = val;
    } else {
        darkMode.value = localStorage.getItem("isLight") == "true";
    }
});
</script>