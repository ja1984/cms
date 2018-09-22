<template>
  <div class="checkbox" :class="{'flip': flip}">
    <input type="checkbox" :checked="shouldBeChecked" :value="value" @change="updateInput" :id="id">
    <label class="checkbox-label" :for="id">
      <slot></slot>
    </label>
  </div>

</template>
<script>
export default {
  name: 'Checkbox',
  data() {
    return {
      test: this.model,
      id: Math.random()
        .toString(36)
        .slice(-5),
    };
  },
  model: {
    prop: 'modelValue',
    event: 'change',
  },
  props: {
    value: {
      type: String,
    },
    flip: {
      type: Boolean,
      default: false,
    },
    modelValue: {
      default: false,
    },
    trueValue: {
      default: true,
    },
    falseValue: {
      default: false,
    },
  },
  computed: {
    shouldBeChecked() {
      if (this.modelValue instanceof Array) {
        return this.modelValue.includes(this.value);
      }
      // Note that `true-value` and `false-value` are camelCase in the JS
      return this.modelValue === this.trueValue;
    },
  },
  methods: {
    updateInput(event) {
      const isChecked = event.target.checked;

      if (this.modelValue instanceof Array) {
        const newValue = [...this.modelValue];

        if (isChecked) {
          newValue.push(this.value);
        } else {
          newValue.splice(newValue.indexOf(this.value), 1);
        }

        this.$emit('change', newValue);
      } else {
        this.$emit('change', isChecked ? this.trueValue : this.falseValue);
      }
    },
  },
};

</script>
<style lang="less" scoped>
@import (reference) '~@/styles/site.less';
@checkbox-checked-color: #acacac;
@checkbox-background-color: #fff;
@checkbox-border-color: #dfdfdf;
@checkbox-size: 2rem;
@checkbox-checked-size: 1rem;
@checkbox-ripple-size: 1.5rem;
.checkbox {
  display: inline-block;
}

*,
*:before,
*:after {
  box-sizing: border-box;
}

input[type='checkbox'] {
  display: none;
  &:checked + label.checkbox-label:before {
    border-color: @checkbox-checked-color;
    animation: ripple 0.2s linear forwards;
  }
  &:checked + label.checkbox-label:after {
    transform: scale(1);
  }
}

label.checkbox-label {
  position: relative;
  padding: .2rem 0 .2rem (@checkbox-size + .5rem);
  margin-bottom: 0;
  background: transparent;
  cursor: pointer;
  vertical-align: bottom;
  &:before,
  &:after {
    position: absolute;
    content: '';
    border-radius: .3rem;
    transition: all 0.3s ease;
    transition-property: transform, border-color;
  }
  &:before {
    left: 0;
    top: 0;
    width: @checkbox-size;
    height: @checkbox-size;
    border: .1rem solid @checkbox-border-color;
    background: @checkbox-background-color;
  }
  &:after {
    top: 0;
    left: 0;
    width: @checkbox-size;
    height: @checkbox-size;
    transform: scale(0);
    background: url(~@/assets/check.svg);
    background-size: 1.3rem;
    background-position: center;
    background-repeat: no-repeat;
  }
}

.checkbox {
  &.flip {
    label {
      padding-left: 0;
      padding-right: (@checkbox-size + 5px);
      &:before {
        left: auto;
        right: 0;
      }
      &:after {
        left: auto;
        right: @checkbox-size / 2 - @checkbox-checked-size / 2;
      }
    }
  }
}
</style>
