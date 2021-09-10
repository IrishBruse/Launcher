<template>
    <div class="h-12 relative flex bg-primary-dark">
        <div class="justify-start self-center px-4">
            <label class="mr-2"> Version: </label>

            <select class="transition duration-300 cursor-pointer bg-primary hover:bg-primary-hover h-8 px-2 outline-none rounded-md" v-model="selectedVersion">
                <option v-for="(version, index) in props.CurrentGame.Versions" :key="index">{{ version.Version }}</option>
            </select>
        </div>
        <button
            class="
                transition
                duration-300
                left-1/2
                absolute
                transform
                bg-tertiary
                text-primary-light
                shadow-xl
                hover:bg-secondary
                -translate-x-1/2 -translate-y-6
                hover:text-primary-light
                rounded-lg
            "
            @click="ButtonClicked"
            ref="playButton"
        >
            <h2 class="text-2xl font-bold px-4 py-2 text-primary-light">{{ buttonText }}</h2>
        </button>
        <div class="transition duration-300 flex-grow flex justify-center hide" ref="progressBar">
            <div class="self-center relative rounded-lg w-11/12 h-6 bg-primary-hover overflow-hidden">
                <div class="absolute h-full bg-tertiary transition-all ease-out" ref="progressBarFill"></div>
                <h2 class="absolute left-1/2 shadow-lg text transform -translate-x-1/2 text-secondary font-bold textOutline">{{ GameDownloadPercent }}%</h2>
            </div>
        </div>
    </div>
</template>

<script setup>
import { onMounted, ref, watch } from "@vue/runtime-core";

const props = defineProps({ CurrentGame: Object });

const GameDownloadPercent = ref(0);

const progressBar = ref();
const progressBarFill = ref();
const playButton = ref();
const selectedVersion = ref();
const buttonText = ref("Download");

watch(
    () => props.CurrentGame.Versions,
    (newVal, oldVal) => {
        selectedVersion.value = newVal[0].Version;
    },
);

watch(
    () => selectedVersion.value,
    (newVal, oldVal) => {
        console.log(newVal);
        console.log(props.CurrentGame.Name);
        console.log(DownloadedVersions[props.CurrentGame.Name]);
        buttonText.value = "Download";
        if (DownloadedVersions[props.CurrentGame.Name] != null) {
            if (DownloadedVersions[props.CurrentGame.Name].includes(newVal)) {
                buttonText.value = "Play";
            }
        }
    },
);

const ButtonClicked = () => {
    if (DownloadedVersions[props.CurrentGame.Name] != null && DownloadedVersions[props.CurrentGame.Name].includes(selectedVersion.value)) {
        buttonText.value = "Play";
        Play(props.CurrentGame.Name + "/" + selectedVersion.value + "/");
    } else {
        TogglePlayButton();
        let downloadUrl;

        props.CurrentGame.Versions.forEach((v) => {
            if (v.Version == selectedVersion.value) {
                downloadUrl = v.Url;
            }
        });

        Download(downloadUrl);
    }
};

onMounted(() => {
    window.addEventListener("DownloadProgressEvent", () => {
        GameDownloadPercent.value = DownloadProgress;
        progressBarFill.value.style.width = DownloadProgress + "%";

        if (DownloadProgress == 100) {
            buttonText.value = "Play";
            TogglePlayButton();
            setTimeout(() => {
                progressBarFill.value.style.width = 0 + "%";
            }, 200);
        }
    });
});

function TogglePlayButton() {
    progressBar.value.classList.toggle("hide");
    playButton.value.classList.toggle("hide");

    setTimeout(() => {
        playButton.value.classList.toggle("disable");
    }, 200);
}
</script>

<style>
.textOutline {
    text-shadow: black 2px 2px 2px;
}

.hide {
    opacity: 0;
    transform: scaleX(0);
}

.disable {
    display: none;
}
</style>
