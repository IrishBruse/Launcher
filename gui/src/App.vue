<template>
    <div class="w-48 flex flex-col">
        <div class="h-16 bg-primary-dark flex justify-center">
            <ThemeToggle></ThemeToggle>
        </div>
        <div class="flex-grow bg-primary"></div>
        <div class="h-12 flex pl-2 bg-primary-dark">
            <button class="self-center flex hover:text-secondary">
                <div class="icons text-2xl hover:text-secondary">settings</div>
                <h3 class="self-center hover:text-secondary">Settings</h3>
            </button>
        </div>
    </div>
    <div class="w-full h-screen flex flex-col self-center justify-center">
        <div class="h-16 bg-primary-dark flex justify-center">
            <h1 class="text-xl self-center text-center">HyperHop</h1>
        </div>
        <div class="flex-grow bg-primary-light"></div>
        <div class="h-12 flex bg-primary-dark">
            <div class="self-center px-4">
                <label for="cars">Version:</label>
                <select class="cursor-pointer" name="cars">
                    <option>0.1.0</option>
                    <option>0.2.0</option>
                    <option>0.3.0</option>
                    <option>0.4.0</option>
                </select>
            </div>
            <button class="absolute left-1/2 border-2 border-solid p-2 c-green">Download</button>
        </div>
    </div>
</template>

<script setup>
import ThemeToggle from "./components/ThemeToggle.vue";

let t = $ref()

t = 123;
</script>
