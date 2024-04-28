<template>
    <div class="playlist">
        <h1>{{ title }}</h1>
        <p>{{ likes }} Like<span v-if="likes > 1">s</span></p>
        <p>{{ description }}</p>
        <div>
            <button>Play</button>
            <button>Shuffle</button>
            <button>Download</button>
            <button>Add to a playlist</button>
            <button>Add to waitlist</button>
            <button>Share</button>
        </div>

        <MusicMiniature v-for="music in musics" :music_id="music.id" :title="music.title" :artist="music.artist" :album="music.album" :album_id="music.album_id" :path="music.path" :illustration="music.illustration"/>
    </div>
</template>

import MusicMiniature from '@/components/music/MusicMiniature.vue'

<script>
export default {
    name: 'Playlist',
    props: {
        playlist_id: Number,
        title: String,
        likes: Number,
        illustration: String,
        description: String,
        musics: Array
    }
    methods: {
        
    }
}

</script>
