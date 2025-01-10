<script lang="ts">
    import { onMount } from "svelte"
    import {ZoomInSolid} from "flowbite-svelte-icons";
    import Zoomist from "zoomist";

    export let image: string;
    let container: HTMLDivElement;
    let currentPercentage: number;

    onMount(() => {
        new Zoomist('.zoomist-container', {
            // Optional parameters
            maxScale: 4,
            bounds: true,
            // if you need slider
            slider: true,
            // if you need zoomer
            zoomer: true
        })
    });
</script>
<!-- zoomist-container -->
<div class="zoomist-container">
    <!-- zoomist-wrapper is required -->
    <div class="zoomist-wrapper">
        <!-- zoomist-image is required -->
        <div class="zoomist-image w-full h-full">
            <!-- you can add anything you want to zoom here. -->
            <img src={image} />
        </div>
    </div>
</div>