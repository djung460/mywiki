var Root = React.createClass({
  getInitialState: function() {
    return {
      articleTitles: [],
      formArticleTitle: ""
    };
  },
  componentDidMount: function() {
    $.getJSON("/api/articles", function(data) {
      this.setState({articleTitles: data["articles"]});
    }.bind(this));
  },
  handleChange: function(evt) {
    evt.preventDefault();
    this.setState({formArticleTitle: evt.target.value});
  },
  handleSubmit: function(e) {
    e.preventDefault();
    $.ajax({
      url: "/api/article",
      data: JSON.stringify({"title": this.state.formArticleTitle}),
      method: "PUT",
      success: function() {
        var newTitles = this.state.articleTitles;
        newTitles.push(this.state.formArticleTitle);
        this.setState({
          formArticleTitle: "",
          articleTitles: newTitles
        });
      }.bind(this)
    });
  },
  render: function() {
    var listElts = [];
    for(var i = 0; i < this.state.articleTitles.length; i++) {
      var name = this.state.articleTitles[i];
      listElts.push(<li key={title} className="list-group-item">{title}</li>);
    }
    return (
      <div className="container articles">
        <div className="row">
          <ul className="list-group">
            {listElts}
          </ul>
        </div>
        <div className="row">
          <h3>Create New Article</h3>
          <form onSubmit={this.handleSubmit}>
            <div className="form-group">
              <label for="articleTitle">Article Title</label>
              <input type="text" className="form-control" value={this.state.formArticleTitle} id="articleTitle" placeholder="Article Title" onChange={this.handleChange}/>
            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
})

ReactDOM.render(<Root/>, document.getElementById("page-content"));
