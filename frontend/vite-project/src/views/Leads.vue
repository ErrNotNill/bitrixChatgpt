<template>
  <div>
    <Sidebar />

    <main id="Home-page">
      <h1>Сделки</h1>
      <div>
        <label for="filter">Filter by:</label>
        <select id="filter" v-model="selectedFilter">
          <option value="all">All</option>
          <option value="Евгений">Евгений</option>
        </select>
        <button @click="applyFilter">Apply Filter</button>
      </div>

      <div class="table-container">
        <ul class="table">
          <li v-for="deal in filteredItems" :key="deal.ID" class="list-item">
            <div class="button-container">
              <button @click="toggleMenu(deal.ID, 'general')" class="table-button">
                {{ deal.TITLE }}
                <p>ID сделки: ({{ deal.ID }})</p>
                Стадия сделки: {{ deal.STAGE_ID }}
              </button>
              <button @click="toggleMenu(deal.ID, 'documents')" class="detail-button">Documents</button>
              <button @click="toggleMenu(deal.ID, 'commentary')" class="detail-button">Commentaries</button>
              <button @click="toggleMenu(deal.ID, 'description')" class="detail-button">Description</button>
            </div>

            <div v-if="activeItem === deal.ID">
              <div v-if="activeSection[deal.ID] === 'documents'" class="item-details">
                <p>Documents for Deal ID: {{ deal.ID }}</p>
                <ul v-if="documentsData[deal.ID]">
                  <li v-for="doc in documentsData[deal.ID]" :key="doc.id">
                    {{ doc.title }}
                    <a :href="doc.downloadUrl" target="_blank">Download</a>
                    <a :href="doc.pdfUrl" target="_blank">PDF</a>
                    <a :href="doc.imageUrl" target="_blank">Image</a>
                  </li>
                </ul>
              </div>

              <div v-if="activeSection[deal.ID] === 'commentary'" class="item-details">
                <p>Commentaries for Deal ID: {{ deal.ID }}</p>
                <ul v-if="commentaryData[deal.ID]">
                  <li v-for="comment in commentaryData[deal.ID]" :key="comment.id">
                    {{ comment.ID }}: {{ comment.COMMENT }}
                  </li>
                </ul>
              </div>

              <div v-if="activeSection[deal.ID] === 'description'" class="item-details">
                <p>Description for Deal ID: {{ deal.ID }}</p>
                <ul v-if="descriptionData[deal.ID]">
                  <li>
                    <!-- Display Description Data -->
                    <p>Title: {{ descriptionData[deal.ID].TITLE }}</p>
                    <!-- Add more fields as necessary -->
                  </li>
                </ul>
              </div>
            </div>
          </li>
        </ul>
      </div>
    </main>
  </div>
</template>

<script>
import axios from 'axios';
import Sidebar from './Sidebar.vue';

export default {
  components: {
    Sidebar,
  },
  data() {
    return {
      jsonArray: [],
      activeItem: null,
      selectedFilter: 'all',
      documentsData: {},
      commentaryData: {},
      descriptionData: {},
      activeSection: {},
    };
  },
  computed: {
    filteredItems() {
      return this.selectedFilter === 'all'
          ? this.jsonArray
          : this.jsonArray.filter((deal) => deal.ASSIGNED_BY_ID === this.selectedFilter);
    },
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      // Your API call to fetch deals
    },
    toggleMenu(dealID, section) {
      if (this.activeItem === dealID && this.activeSection[dealID] === section) {
        this.activeItem = null;
        this.activeSection[dealID] = null;
      } else {
        this.activeItem = dealID;
        this.activeSection[dealID] = section;
        // Call the appropriate function based on the section
        if (section === 'documents') this.showDocuments(dealID);
        if (section === 'commentary') this.showCommentary(dealID);
        if (section === 'description') this.showDescription(dealID);
      }
    },
    showDocuments(dealID) {
      // Fetch and assign documents data
    },
    showCommentary(dealID) {
      // Fetch and assign commentary data
    },
    showDescription(dealID) {
      // Fetch and assign description data
    },
    applyFilter() {
      // Filter application logic
    },
  },
};
</script>


<style>

/* Style for the table button */
.table-button {
  width: 300px;
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
  width: 500px;
  height: 300px;
  margin-top: 10px;
  background-color: lightgray;
  padding: 5px;
  border-radius: 5px;
}

/* Remove bullet points from list items */
.list-item {
  width: 100%;
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
  text-align: left;
  width: 100%;
  margin-top: 40px;
}
.table li {
  float: left;
}
.detail-button {
  background-color: #f0f0f0; /* Light grey, customizable */
  border: 1px solid #d0d0d0; /* Slightly darker grey border */
  color: #333; /* Dark grey text */
  padding: 5px 15px;
  margin: 5px;
  cursor: pointer;
  transition: background-color 0.3s ease; /* Smooth transition for hover effect */
}

.detail-button:hover {
  background-color: #e0e0e0; /* Slightly darker grey on hover */
}



</style>