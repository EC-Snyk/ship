import React from "react";
import ConfigItemTitle from "./ConfigItemTitle";

export default class ConfigTextarea extends React.Component {

  constructor(props) {
    super(props)
    this.textareaRef = React.createRef();
    this.state = {
      textareaVal: "",
      focused: false
    }
  }

  handleOnChange = (e) => {
    const { handleOnChange, name } = this.props;
    this.setState({ textareaVal: e.target.value });
    if (handleOnChange && typeof handleOnChange === "function") {
      handleOnChange(name, e.target.value);
    }
  }

  componentDidUpdate(lastProps) {
    if (this.props.value !== lastProps.value && !this.state.focused) {
      this.setState({ textareaVal: this.props.value });
    }
  }

  componentDidMount() {
    if (this.props.value) {
      this.setState({ textareaVal: this.props.value });
    }
  }

  render() {
    var hidden = this.props.hidden || this.props.when === "false";

    return (
      <div className={`field field-type-text u-marginTop--15 ${hidden ? "hidden" : ""}`}>
        {this.props.title !== "" ?
          <ConfigItemTitle
            title={this.props.title}
            recommended={this.props.recommended}
            required={this.props.required}
            name={this.props.name}
          />
          : null}
        {this.props.help_text !== "" ? <p className="field-section-help-text u-marginTop--small u-lineHeight--normal">{this.props.help_text}</p> : null}
        <div className="field-input-wrapper u-marginTop--15">
          <textarea
            ref={this.textareaRef}
            {...this.props.props}
            placeholder={this.props.default}
            value={this.state.textareaVal}
            readOnly={this.props.readonly}
            disabled={this.props.readonly}
            onChange={(e) => this.handleOnChange(e)}
            onFocus={() => this.setState({ focused: true })}
            onBlur={() => this.setState({ focused: false })}
            className={`${this.props.className || ""} Textarea ${this.props.readonly ? "readonly" : ""}`}>
          </textarea>
        </div>
      </div>
    );
  }
}
