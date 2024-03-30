<template>

    <!-- Sidebar -->
  <Sidebar />

  <main id="Home-page">
    <h1>Сделки</h1>
    <p></p>
    <!-- Filter dropdown -->
    <div>
      <label for="filter">Filter by: </label>
      <select id="filter" v-model="selectedFilter">
        <option value="all">All</option>
        <option value="Евгений">Opts...</option>

        <!-- Add more filter options here as needed -->
      </select>
      <p></p>
      <button @click="applyFilter">Apply Filter</button>
    </div>

    <div class="table-container">
      <ul class="table">
        <li v-for="deal in jsonArray" :key="deal.id" class="list-item">
          <div class="button-container">
            <button @click="toggleMenu(deal.id)" class="table-button">
              {{ deal.title }} ({{ deal.id }})
            </button>
          </div>
          <div v-if="activeItem === deal.id" class="item-details">
            <!-- Adjust according to the actual properties of a deal -->
            <p>Title: {{ deal.title }}</p>
            <!-- Add more details as needed -->
          </div>
        </li>
      </ul>

    </div>
  </main>

</template>

<script>
import axios from 'axios'
import Sidebar from '@/components/Sidebar.vue'

export default {
  components: { Sidebar },
  data() {
    return {
      jsonArray: [],
    }
  },
  created() {
    axios.get('http://localhost:9090/api/deals_get')
        .then((response) => {
          console.log(response.data); // Log the response data to see its structure
          this.jsonArray = response.data.result; // Make sure this path matches the response's structure
        })
        .catch((error) => {
          console.error('Error fetching data:', error);
        });
  },
  computed: {
    filteredItems() {
      if (this.selectedFilter === 'all') {
        return this.jsonArray;
      } else {
        return this.jsonArray.filter((deal) => deal.ASSIGNED_BY_ID === this.selectedFilter);
      }
    }
  },
  methods: {
    toggleMenu(ID) {
      this.activeItem = this.activeItem === ID ? null : ID;
    },
    applyFilter() {
      if (this.selectedFilter === 'Евгений') {
        this.selectedFilter = 'Евгений';
      }
    },
    loadMore() {
      this.itemsToShow += 10; // Increase the number of items to show
    },
  },
};
</script>

<style>
/* Style for the table button */
.table-button {
  background-color: green;
  border: none;
  color: white;
  padding: 5px 10px;
  margin: 5px;
  cursor: pointer;
  text-align: center;
}

/* Style for the item details */
.item-details {
  margin-top: 10px;
  background-color: lightgray;
  padding: 5px;
  border-radius: 5px;
}

/* Remove bullet points from list items */
.list-item {
  list-style-type: none;
  margin-left: 0;
  padding-left: 0;
  background-color: transparent; /* Ensure consistent background */
}

/* Highlighted background color for 'Евгений' */
.highlighted {
  background-color: red;
}

/* Position the table on the left side */
/* Position the table on the left side */
.table-container {
  text-align: center;
  width: 50%;
  margin: 0 auto;
}
.table li {
  float: left;
}
</style>