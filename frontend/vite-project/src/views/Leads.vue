<template>
  <!-- Sidebar -->
  <Sidebar />

  <main id="Home-page">
    <h1>Сделки</h1>
    <div>
      <!-- Filter dropdown -->
      <div>
        <label for="filter">Filter by: </label>
        <select id="filter" v-model="selectedFilter">
          <option value="all">All</option>
          <option value="Евгений">Евгений</option>
          <!-- Add more filter options here as needed -->
        </select>
        <button @click="applyFilter">Apply Filter</button>
      </div>

      <div class="table-container">
        <ul class="table">
          <li v-for="deal in jsonArray" :key="deal.ID" class="list-item">
            <div class="button-container">
              <button @click="toggleMenu(deal.ID)" class="table-button">
                {{ deal.TITLE }}
                <p>ID сделки: ({{ deal.ID }})</p>
                Стадия сделки: {{ deal.STAGE_ID }}
              </button>
            </div>

            <!-- Document Details -->
            <div v-if="activeItem === deal.ID && activeSection[deal.ID] === 'documents'" class="item-details">
              <button @click="showDocuments(deal.ID)" class="detail-button">Documents</button>
              <div v-if="documentsData[deal.ID]">
                <ul>
                  <li v-for="doc in documentsData[deal.ID]" :key="doc.id">
                    {{ doc.title }}
                    <!-- Links -->
                    <a :href="doc.downloadUrl" target="_blank">Download</a>
                    <a :href="doc.pdfUrl" target="_blank">PDF</a>
                    <a :href="doc.imageUrl" target="_blank">Image</a>
                  </li>
                </ul>
              </div>
            </div>

            <!-- Commentary Details -->
            <div v-if="activeItem === deal.ID && activeSection[deal.ID] === 'commentary'" class="item-details">
              <button @click="showCommentary(deal.ID)" class="detail-button">Commentaries</button>
              <div v-if="commentaryData[deal.ID]">
                <ul>
                  <li v-for="com in commentaryData[deal.ID]" :key="com.id">
                    {{ com.ID }}: {{ com.COMMENT }}
                  </li>
                </ul>
              </div>
            </div>

            <!-- Description Details -->
            <div v-if="activeItem === deal.ID && activeSection[deal.ID] === 'description'" class="item-details">
              <button @click="showDescription(deal.ID)" class="detail-button">Description</button>
              <div v-if="descriptionData[deal.ID]">
                <!-- Display Description Fields -->
                <p>ID: {{ descriptionData[deal.ID].ID }}</p>
                <p>Title: {{ descriptionData[deal.ID].TITLE }}</p>
                <p>Type ID: {{ descriptionData[deal.ID].TYPE_ID }}</p>
                <p>Stage ID: {{ descriptionData[deal.ID].STAGE_ID }}</p>
                <p>Opportunity: {{ descriptionData[deal.ID].OPPORTUNITY }}</p>
                <p>Currency ID: {{ descriptionData[deal.ID].CURRENCY_ID }}</p>
                <p>Begindate: {{ descriptionData[deal.ID].BEGINDATE }}</p>
                <p>Closedate: {{ descriptionData[deal.ID].CLOSEDATE }}</p>
                <p>Assigned_by_id: {{ descriptionData[deal.ID].ASSIGNED_BY_ID }}</p>
                <p>Created+by_id: {{ descriptionData[deal.ID].CREATED_BY_ID }}</p>
                <p>Modify_by_id: {{ descriptionData[deal.ID].MODIFY_BY_ID }}</p>
                <p>Date_create: {{ descriptionData[deal.ID].DATE_CREATE }}</p>
                <p>Date_modify: {{ descriptionData[deal.ID].DATE_MODIFY }}</p>
                <!-- Add more fields as needed -->
              </div>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </main>
</template>

<script>
import axios from 'axios';
import Sidebar from '@/components/Sidebar.vue';

export default {
  components: { Sidebar },
  data() {
    return {
      jsonArray: [],
      activeItems: {},
      selectedFilter: 'all',
      documentsData: {},
      commentaryData: {},
      descriptionData: {},
    };
  },
  created() {
    this.fetchDeals();
  },
  computed: {
    filteredItems() {
      return this.selectedFilter === 'all'
          ? this.jsonArray
          : this.jsonArray.filter(deal => deal.ASSIGNED_BY_ID === this.selectedFilter);
    },
  },
  methods: {
    fetchDeals() {
      // API call to fetch deals
    },
    toggleSection(dealID, section) {
      if (this.activeItems[dealID] === section) {
        this.$set(this.activeItems, dealID, null); // Close the section
      } else {
        this.$set(this.activeItems, dealID, section); // Open the section
        this.loadData(dealID, section); // Fetch data for the section
      }
    },
    isActive(dealID, section) {
      return this.activeItems[dealID] === section;
    },
    applyFilter() {
      // Implementation for applying selected filter
    },
    loadData(dealID, section) {
      switch (section) {
        case 'documents':
          this.showDocuments(dealID);
          break;
        case 'commentary':
          this.showCommentary(dealID);
          break;
        case 'description':
          this.showDescription(dealID);
          break;
      }
    },
    showDocuments(dealID) {
      axios.get(`https://b24app.rwp2.com/api/documents/${dealID}`)
          .then(response => {
            this.$set(this.documentsData, dealID, response.data.result.documents);
          })
          .catch(error => {
            console.error('Error fetching documents:', error);
          });
    },
    showCommentary(dealID) {
      axios.get(`https://b24app.rwp2.com/api/comments/${dealID}`)
          .then(response => {
            this.$set(this.commentaryData, dealID, response.data.result);
          })
          .catch(error => {
            console.error('Error fetching commentary:', error);
          });
    },
    showDescription(dealID) {
      axios.get(`https://b24app.rwp2.com/api/description/${dealID}`)
          .then(response => {
            this.$set(this.descriptionData, dealID, response.data.result);
          })
          .catch(error => {
            console.error('Error fetching description:', error);
          });
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