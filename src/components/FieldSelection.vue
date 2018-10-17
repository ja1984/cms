<template>
  <ul>
    <li>
      <button class="toolbar-button" @click="addField('input')">
        <i class="fas fa-toggle-on fa-fw"></i> Text</button>
    </li>
    <li>
      <button class="toolbar-button" @click="addField('textarea')">
        <i class="fas fa-code fa-fw"></i> Textarea</button>
    </li>
    <li>
      <button class="toolbar-button" @click="addField('boolean')">
        <i class="fas fa-toggle-on fa-fw"></i> True / false</button>
    </li>
    <li>
      <button class="toolbar-button" @click="addField('number')">
        <i class="fas fa-hashtag fa-fw"></i> Number</button>
    </li>
    <li>
      <button class="toolbar-button" @click="addField('date')">
        <i class="far fa-calendar-alt fa-fw"></i> Date</button>
    </li>
    <!-- <li>
                  <button class="toolbar-button" @click="addField('group')">
                    <i class="far fa-folder fa-fw"></i> Group</button>
                </li> -->
    <li>
      <button class="toolbar-button" @click="addField('media')">
        <i class="far fa-image fa-fw"></i> Media</button>
    </li>
    <li>
      <button class="toolbar-button" @click="addField('select')">
        <i class="far fa-list-alt fa-fw"></i> List</button>
    </li>
    <li>
      <button :disabled="disableRepeater" class="toolbar-button" @click="addField('repeater')">
        <i class="fas fa-redo"></i> Repeater</button>
    </li>
    <li></li>
    <li></li>
  </ul>
</template>

<script>
export default {
  name: 'FieldSelection',
  props: {
    disableRepeater: { type: Boolean, default: false },
  },
  methods: {
    addField(id) {
      this.$emit('addField', id);
    },
  },
};
</script>

<style lang="less" scoped>
li {
  i {
    font-size: 2rem;
    vertical-align: middle;
    margin-right: 0.5rem;
  }
}

.toolbar-button {
  border: none;
  padding: 0.75rem;
  cursor: pointer;
  transition: all ease 0.3s;
  // font-weight: 600;
  font-size: 1.5rem;
  display: block;
  width: 100%;
  text-align: left;
  background: transparent;

  &:active,
  &:focus {
    outline: none;
  }

  &:hover {
    background: #f9f9f9;
    color: #282d3a;
  }
}
</style>
