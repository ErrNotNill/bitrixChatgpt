<template>
  <Sidebar />
  <main id="settings-page">
    <h1>Settings</h1>
    <p>User fields</p>
    <div class="block-container" v-for="(block, index) in blocks" :key="index" style="margin-bottom: 5px;">

      <div class="block_uf">
        <input
            v-model="block.input_field1"
            class="input-text"
            placeholder="Input User_Field"
        />
        <input
            v-model="block.input_field2"
            class="input-text"
            placeholder="Input Title"
        />
        <button @click="sendData(block)" style="color: #265df2">Send</button>
      </div>
      <div class="block_info">
        {{ block_info[index] }}
      </div>
    </div>
    <button @click="saveData" style="color: #265df2; margin-top: 20px;">Save</button>
  </main>
</template>

<script>
import Sidebar from '@/components/Sidebar.vue'
export default {
  components: { Sidebar },
  data() {
    return {
      blocks: Array(5).fill().map(() => ({
        input_field1: '',
        input_field2: ''
      })),
      block_info: Array(5).fill(''),
      data_to_serve: [] // Array to collect all info_blocks data
    }
  },
  methods: {
    sendData(block) {
      const dataToSend = JSON.stringify({
        input_field1: block.input_field1,
        input_field2: block.input_field2
      });
      console.log(dataToSend);
      // Here you would typically send the data to your server
      // For the purpose of this example, we'll just simulate updating the block_info
      const index = this.blocks.indexOf(block);
      this.block_info[index] = dataToSend;
      // Additionally, collect the data in the data_to_serve array
      this.data_to_serve.push({input_field1: block.input_field1, input_field2: block.input_field2});
    },
    saveData() {
      console.log("Saving data to server:", JSON.stringify(this.data_to_serve));
      // Here you would send data_to_serve to your server
      fetch('https://b24app.rwp2.com/api/save_settings', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(this.data_to_serve),
      })
          .then(response => response.json())
          .then(data => console.log('Success:', data))
          .catch((error) => console.error('Error:', error));
    }
  }
}
</script>

<style scoped>
.block-container {
  display: flex;
  align-items: center;
}

.block_uf {
  background-color: lightgrey;
  width: 30%;
  height: 50px;
  padding: 5px;
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.input-text {
  background-color: white;
  color: black;
  margin-right: 5px;
  padding: 5px;
}

.block_info {
  background-color: lightgray;
  width: 35%;
  height: 50px;
  margin-left: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
