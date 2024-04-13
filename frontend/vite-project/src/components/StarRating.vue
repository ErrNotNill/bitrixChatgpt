<template>
  <div class="rating-area">
    <input type="radio" id="star-10" name="rating" value="10">
    <label for="star-5" title="Оценка «10»"></label>
    <input type="radio" id="star-9" name="rating" value="9">
    <label for="star-4" title="Оценка «9»"></label>
    <input type="radio" id="star-8" name="rating" value="8">
    <label for="star-3" title="Оценка «8»"></label>
    <input type="radio" id="star-7" name="rating" value="7">
    <label for="star-2" title="Оценка «7»"></label>
    <input type="radio" id="star-6" name="rating" value="6">
    <label for="star-1" title="Оценка «6»"></label>
    <input type="radio" id="star-5" name="rating" value="5">
    <label for="star-5" title="Оценка «5»"></label>
    <input type="radio" id="star-4" name="rating" value="4">
    <label for="star-4" title="Оценка «4»"></label>
    <input type="radio" id="star-3" name="rating" value="3">
    <label for="star-3" title="Оценка «3»"></label>
    <input type="radio" id="star-2" name="rating" value="2">
    <label for="star-2" title="Оценка «2»"></label>
    <input type="radio" id="star-1" name="rating" value="1">
    <label for="star-1" title="Оценка «1»"></label>
  </div>
</template>

<script>
export default {
  props: {
    value: {
      type: Number,
      default: 0
    },
    maxStars: {
      type: Number,
      default: 5
    }
  },
  data() {
    return {
      selectedStars: this.value
    };
  },
  methods: {
    selectRating(index) {
      this.selectedStars = index;
      this.$emit('input', this.selectedStars);
    }
  }
};
</script>

<style scoped>
.rating-area {
  overflow: hidden;
  width: 265px;
  margin: 0 auto;
}
.rating-area:not(:checked) > input {
  display: none;
}
.rating-area:not(:checked) > label {
  float: right;
  width: 42px;
  padding: 0;
  cursor: pointer;
  font-size: 32px;
  line-height: 32px;
  color: lightgrey;
  text-shadow: 1px 1px #bbb;
}
.rating-area:not(:checked) > label:before {
  content: '★';
}
.rating-area > input:checked ~ label {
  color: gold;
  text-shadow: 1px 1px #c60;
}
.rating-area:not(:checked) > label:hover,
.rating-area:not(:checked) > label:hover ~ label {
  color: gold;
}
.rating-area > input:checked + label:hover,
.rating-area > input:checked + label:hover ~ label,
.rating-area > input:checked ~ label:hover,
.rating-area > input:checked ~ label:hover ~ label,
.rating-area > label:hover ~ input:checked ~ label {
  color: gold;
  text-shadow: 1px 1px goldenrod;
}
.rate-area > label:active {
  position: relative;
}
</style>