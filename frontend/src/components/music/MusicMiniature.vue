<template>
    <div class="music-miniature" style="display: flex; justify-content: space-between; align-items: center;">
        <div>
            <h1>{{ title }}</h1>
            <p>{{ artist }} <span v-if="album">- {{ album }}</span></p>
        </div>
        <select>
            <option value="download">Download</option>
            <option value="add_to_playlist">Add to a playlist</option>
            <option value="add_to_waitlist">Add to waitlist</option>
            <span v-if="album_id"><option value="show_album">Show album</option></span>
            <option value="show_artist">Show artist</option>
            <option value="share">Share</option>
        </select>
    </div>
</template>

<script>
export default {
    name: 'MusicMiniature',
    props: {
        music_id: Number,
        title: String,
        artist: String,
        album: String,
        album_id: Number,
        path: String
    }
}
</script>
