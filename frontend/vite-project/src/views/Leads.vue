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
        <li v-for="deal in jsonArray" :key="deal.ID" class="list-item">
          <div class="button-container">
            <button @click="toggleMenu(deal.ID)" class="table-button">
              {{ deal.TITLE }}
              <p>ID сделки: ({{ deal.ID }})</p>
              Стадия сделки: {{ deal.STAGE_ID }}
            </button>
          </div>
          <div v-if="activeItem === deal.ID" class="item-details">
            <!-- Adjust according to the actual properties of a deal -->
            <!-- Example detail -->
            <p>Title: {{ deal.TITLE }}</p>

            <!-- Additional Buttons -->
            <button @click="showDocuments(deal.ID)" class="detail-button">Documents</button>
            <!-- Conditional rendering based on whether documentsData for this deal.ID exists -->
            <div v-if="activeItem === deal.ID" :class="{'item-details': true, 'visible': activeContent[deal.ID] === 'documents'}">
              <div v-if="documentsData[deal.ID] && documentsData[deal.ID].length">
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
            <div v-if="activeItem === deal.ID" :class="{'item-details': true, 'visible': activeContent[deal.ID] === 'commentary'}">
            <button @click="showCommentary(deal.ID)" class="detail-button">Commentary</button>
            </div>
            <div v-if="activeItem === deal.ID" :class="{'item-details': true, 'visible': activeContent[deal.ID] === 'description'}">
            <button @click="showDescription(deal.ID)" class="detail-button">Description</button>
            </div>
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
      activeItem: null,
      selectedFilter: 'all', // Default to "All"
      itemsPerPage: 50, // Number of items to show initially
      itemsToShow: 50, // Number of items to show currently
      documentsData: {},
      activeContent: {},
    }
  },
  created() {
    axios.get('https://b24app.rwp2.com/api/deals_get')
        .then((response) => {
          this.jsonArray = response.data.result; // Correct path to the data
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
    showDocuments(ID) {
      if (this.activeContent[ID] !== 'documents') {
        this.activeContent[ID] = 'documents';
        axios.get(`https://b24app.rwp2.com/api/documents/${ID}`)
            .then(response => {
              // Direct assignment for Vue 3 reactivity
              this.documentsData[ID] = response.data.result.documents;
            })
            .catch(error => {
              console.error('Error fetching documents:', error);
            });
      } else {
        this.activeContent[ID] = ''; // Collapse if already open
      }
    },

    showCommentary(ID) {
      this.activeContent[ID] = this.activeContent[ID] !== 'commentary' ? 'commentary' : '';
      // Fetch commentary logic...
    },
    showDescription(ID) {
      this.activeContent[ID] = this.activeContent[ID] !== 'description' ? 'description' : '';
      // Fetch description logic...
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
  transition: max-height 0.3s ease-out, opacity 0.3s ease; /* Smooth transition */
  overflow: hidden; /* Keeps content clipped */
  max-height: 0; /* Start with content collapsed */
  opacity: 0; /* Content is initially invisible */
}

.item-details.visible {
  opacity: 1; /* Make content visible */
  max-height: 500px; /* Enough height to show content */
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

.button-container {
  display: flex;
  justify-content: space-around; /* Adjust spacing as needed */
  margin-bottom: 10px; /* Adds space between buttons and content */
}


</style>