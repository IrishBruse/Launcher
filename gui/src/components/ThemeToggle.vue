<template>
    <div class="h-16 bg-primary-dark flex justify-center">
        <a class="w-10 h-10 self-center select-none cursor-pointer" @click="toggleTheme">
            <span class="icons transition duration-300 hover:text-secondary absolute text-3xl" :class="isLight ? 'opacity-0 z-0' : 'opacity-100 z-10'">
                dark_mode
            </span>
            <span class="icons transition duration-300 hover:text-secondary absolute text-3xl" :class="isLight ? 'opacity-100 z-10' : 'opacity-0 z-0'">
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
    var body = document.getElementsByTagName("body")[0];
    body.classList.toggle("light", isLight.value);
});
</script>