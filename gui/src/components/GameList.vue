
<template>
    <div class="flex-grow bg-primary-light" ref="list">
        <button v-for="project in projectData" :key="project.Name" class="gameSelect" @click="changeProject($event.currentTarget, project)">
            <img :src="project.IconURL" class="h-full transition-all duration-300 rounded-3xl group-hover:rounded-xl bg-white" />
            <h1 class="text-left flex-grow self-center px-2 overflow-hidden overflow-ellipsis">{{ project.Name }}</h1>
        </button>
    </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

const emit = defineEmits(["projectChanged"]);

const projectData = ref();
const list = ref();

const changeProject = (elem, currentProject) => {
    for (let i = 0; i < list.value.children.length; i++) {
        var current = elem == list.value.children[i];
        list.value.children[i].classList.toggle("current", current);
    }

    emit("projectChanged", currentProject.Name, currentProject.Versions);
};

onMounted(() => {
    document.addEventListener("DownloadedProjectsMetadata", () => {
        projectData.value = projectsMetadata;
        setTimeout(() => {
            emit("projectChanged", "Pick a game", ["none"]);
        }, 1);
    });
});
</script>

<style >
.gameSelect {
    @apply group transition duration-300 w-full h-14 relative flex p-2 cursor-pointer hover:bg-primary-hover;
}

.gameSelect.current > div {
    @apply scale-y-100;
}

.gameSelect.current > img {
    @apply left-1 rounded-xl;
}

.gameSelect.current > h1 {
    @apply font-bold;
}

.clipLoadingCircle {
    clip-path: polygon(0% 0%, 100% 0%, 100% 50%, 0% 50%);
}
</style>