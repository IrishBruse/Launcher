<template>
    <div class="h-16 bg-primary-light flex justify-center">
        <a class="absolute w-10 h-10 self-center select-none cursor-pointer" @click="toggleTheme">
            <span class="icons hover:text-secondary absolute text-3xl transition duration-400" :class="isLight ? 'opacity-0 z-0' : 'opacity-100 z-10'">
                dark_mode
            </span>
            <span class="icons hover:text-secondary absolute text-3xl transition duration-400" :class="isLight ? 'opacity-100 z-10' : 'opacity-0 z-0'">
                light_mode
            </span>
        </a>
    </div>
</template>

<script setup>
import { ref, onMounted, defineEmits } from "vue";

const isLight = ref(false);

const emit = defineEmits(["toggledTheme"]);

const toggleTheme = () => {
    isLight.value = !isLight.value;
    localStorage.setItem("isLight", isLight.value);
    var body = document.getElementsByTagName("body")[0];
    body.classList.toggle("light", isLight.value);
    emit("toggledTheme");
};

onMounted(() => {
    if (localStorage.getItem("isLight") === null) {
        isLight.value = !window.matchMedia("(prefers-color-scheme: dark)").matches;
        localStorage.setItem("isLight", isLight.value);
    } else {
        isLight.value = localStorage.getItem("isLight") == "true";
    }

    var body = document.getElementsByTagName("body")[0];
    body.classList.toggle("light", isLight.value);
});
</script>