<!-- Developed by Taipei Urban Intelligence Center 2023-2024-->

<script setup>
import { computed, ref } from "vue";
import { useDialogStore } from "../../store/dialogStore";
import { useMapStore } from "../../store/mapStore";
import DialogContainer from "./DialogContainer.vue";
const tempGeoJson = ref({
	type: "Feature",
	properties: {
		marker_name: "",
		marker_memo: "",
	},
	geometry: { type: "Point", coordinates: null },
});
const dialogStore = useDialogStore();
const mapStore = useMapStore();
const params = ref({
	name: "",
	longitude: null,
	latitude: null,
});
const viewPoint = ref({
	name: "",
	coordinates: [],
	zoom: null,
	pitch: null,
	bearing: null,
});

function handleClose() {
	dialogStore.hideAllDialogs();
	params.value.longitude = null;
	params.value.latitude = null;
}

const pinName = ref("");

const handleAddPin = () => {
	if (!pinName.value) {
		dialogStore.showNotification("fail", "請輸入地標名稱");
		return;
	}
	// const { lng, lat } = mapStore.map.getCenter();
	// const zoom = mapStore.map.getZoom();
	// const pitch = mapStore.map.getPitch();
	// const bearing = mapStore.map.getBearing();
	// viewPoint.value.zoom = zoom;
	// viewPoint.value.pitch = pitch;
	// viewPoint.value.bearing = bearing;
	// viewPoint.value.coordinates = [lng, lat];
	// const viewPointArray = [
	// 	[lng, lat],
	// 	zoom,
	// 	pitch,
	// 	bearing,
	// 	viewPoint.value.name,
	// ];
	// mapStore.addViewPoint(viewPointArray);
	mapStore.addMarker(pinName.value);
	dialogStore.hideAllDialogs();
	dialogStore.showNotification("success", "新增地標成功");
	// viewPoint.value = {
	// 	name: "",
	// 	coordinates: [],
	// 	zoom: null,
	// 	pitch: null,
	// 	bearing: null,
	// };
	// mapStore.easeToLocation(viewPointArray);
};
</script>

<template>
	<DialogContainer dialog="addPin" @on-close="handleClose">
		<div class="login add-mark-to-map">
			<h1 class="title">建立地標</h1>

			<div class="content">
				<label for="view-point-name">地標名稱：</label>
				<input
					v-model="pinName"
					type="text"
					id="view-point-name"
					name="view-point-name"
					placeholder="請輸入地標名稱"
					required
				/>
				<!-- 是否要建立地標？ -->
			</div>
			<!-- <label for="marker-name">地標名稱：</label>
			<input
				v-model="params.name"
				type="text"
				id="marker-name"
				name="marker-name"
				placeholder="請輸入地標名稱"
				required
			/>
			<label for="longitude">經度：</label>
			<input
				v-model="params.longitude"
				type="number"
				id="longitude"
				name="longitude"
				min="-180"
				max="180"
				step="any"
				placeholder="請輸入有效的經度，範圍在-180到180之間"
				required
			/>

			<label for="latitude">緯度：</label>
			<input
				v-model="params.latitude"
				type="number"
				id="latitude"
				name="latitude"
				min="-90"
				max="90"
				step="any"
				placeholder="請輸入有效的緯度，範圍在-90到90之間"
				required
			/> -->
			<!-- <div class="content">是否儲存當前視角？</div> -->

			<div class="btn-box">
				<button @click.prevent="handleAddPin">確認</button>
			</div>
		</div>
	</DialogContainer>
</template>

<style scoped lang="scss">
.login {
	width: 300px;

	p {
		text-align: center;
		color: var(--color-complement-text);
	}

	// p:last-child {
	// 	background: linear-gradient(
	// 		75deg,
	// 		var(--color-complement-text),
	// 		var(--color-highlight) 70%,
	// 		var(--color-complement-text)
	// 	);
	// 	background-clip: text;
	// 	-webkit-background-clip: text;
	// 	-webkit-text-fill-color: transparent;
	// 	animation: title-gradient 10s linear none infinite;
	// }

	a {
		color: var(--color-highlight);
	}

	label {
		margin-bottom: 4px;
		color: var(--color-complement-text);
		font-size: var(--font-s);
		align-self: flex-start;
	}

	input {
		margin-bottom: 8px;
		width: calc(100% - 14px);
	}

	&-logo {
		display: flex;
		justify-content: center;

		h1 {
			font-weight: 500;
		}

		h2 {
			font-size: var(--font-s);
			font-weight: 400;
		}

		&-image {
			width: 22.94px;
			height: 45px;
			margin: 0 10px 0 0;

			img {
				height: 45px;
				filter: invert(1);
			}
		}
	}

	&-form {
		width: 100%;
		height: 200px;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
}
.add-mark-to-map {
	// text-align: center;
	label {
		text-align: center;
	}
	.content {
		margin: 30px 0 30px;
	}
	h1 {
		text-align: center;
	}
	input {
		font-size: var(--font-s);
		margin-bottom: 4px;
		margin-top: 4px;
	}
	.btn-box {
		display: flex;
		align-items: center;
		justify-content: center;
	}
	button {
		width: 80px;
		display: flex;
		align-items: center;
		justify-content: center;
		margin: 12px 0;
		padding: 6px;
		font-size: var(--font-s);
		background-color: #03b2c3;
		border-radius: 100px;

		img {
			width: 1.5rem;
			margin: 0 10px 0 0;
		}
	}
	/* Chrome, Safari, Edge, Opera */
	input::-webkit-outer-spin-button,
	input::-webkit-inner-spin-button {
		-webkit-appearance: none;
		margin: 0;
	}

	/* Firefox */
	input[type="number"] {
		-moz-appearance: textfield;
	}
}
</style>
