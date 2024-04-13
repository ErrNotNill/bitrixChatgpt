<template>
  <div class="container">
    <div class="page-container">
      <img src="../../public/harizma.jpg" alt="Image Description" class="logo">
    </div>
    <div class="feedback-form-container">
      <h2 class="form-heading">Оцените Харизму</h2>
      <form @submit.prevent="submitFeedback">
        <div class="form-field">
          <label for="rating">Ваша оценка (1-10): *</label>
          <select v-model="rating" required class="rating-select">
            <option value="0">Выберите вариант</option>
            <option v-for="num in 10" :value="num" :key="num">{{ num }}</option>
          </select>
          <div class="mandatory-note">оценка обязательна</div>
        </div>
        <div class="form-field">
          <label for="comment">Что нам улучшить?</label>
          <textarea id="comment" v-model="comment"></textarea> <!-- Added id="comment" -->
        </div>
        <div class="form-field">
          <input type="submit" value="Отправить">
        </div>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      rating: 0,
      comment: ''
    }
  },
  methods: {
    submitFeedback() {
      const feedbackData = {
        rating: this.rating,
        comment: this.comment
      };
      console.log(feedbackData);

      fetch('https://harizma-service.ru/api/user-form', {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify(feedbackData)
      })
          .then(() => {
            this.$emit('formSubmitted');
          })
          .catch(error => {
            console.error('Error:', error);
          });
    }
  }
}
</script>

<style scoped>

.rating-select {
  height: 50px;
}

.container {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}

.page-container {
  margin-bottom: 20px; /* Adjust spacing between containers */
}

.logo {
  width: 1024px;
  height: 580px;
}

.feedback-form-container {
  width: 300px;
  box-sizing: border-box;
}

.form-heading {
  color: #FFB500;
  text-align: center;
  margin-bottom: 30px;
}

.form-field {
  margin-bottom: 20px;
}

.form-field label {
  display: block;
  width: 100%;
  color: white;
}

.mandatory-note {
  color: red;
}

.form-field select, .form-field textarea {
  display: block;
  width: 100%;
  color: white;
  background-color: #333;
  border: none;
}

.form-field input[type="submit"] {
  width: 300px;
  height: 40px;
  color: white;
  background-color: #333;
  border: 1px solid #D3D3D3;
}

.form-field input[type="submit"] {
  padding: 0;
  cursor: pointer;
  background-color: #FFB500;
}

.form-field input[type="submit"]:hover {
  background-color: #FF7F50;
}

#comment {
  width: 100%;
  height: 80px; /* Set the height to 100px */
}
</style>
