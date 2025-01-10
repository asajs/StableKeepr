<script lang="ts">
    import { onMount } from "svelte"
    import Zoomist from "zoomist";

    export let image: string;
    let imageElement: HTMLDivElement;

    onMount(() => {
        imageElement.onload = () => {
            new Zoomist('.zoomist-container', {
                maxScale: 8,
                bounds: true,
                slider: true,
                zoomer: true
            })
        };
    });
</script>
<div class="zoomist-container">
    <div class="zoomist-wrapper">
        <div class="zoomist-image w-full h-full">
            <img bind:this={imageElement} src={image} class="w-full h-full" />
        </div>
    </div>
</div>