var Root = React.createClass({
  getInitialState: function() {
    return {
      articleTitles: [],
    };
  },
  componentDidMount: function() {
    $.getJSON("/api/articles", function(data) {
      var articles = data["articles"];
      var titles = [];
      for(var i = 0; i < articles.length; i++) {
        var article = JSON.parse(articles[i]);
        titles.push(article["title"]);
      }
      this.setState({articleTitles: titles});
      console.dir(data["articles"]);
    }.bind(this));
  },
  handleChange: function(evt) {
    evt.preventDefault();

    const target = evt.target
    const val = target.value
    const name = target.name

    // For handling multiple inputs
    this.setState({
      [name]: val
    });
  },
  handleSubmit: function(e) {
    e.preventDefault();
    $.ajax({
      url: "/api/article",
      data: JSON.stringify({"title":this.state.formArticleTitle, "content": this.state.formArticleContent}),
      method:'PUT',
      type:'POST',
      success: function() {
        var newTitles = this.state.articleTitles;
        newTitles.push(this.state.formArticleTitle);
        this.setState({
          formArticleTitle: "",
          formArticleContent: "",
          articleTitles: newTitles,
        });
      }.bind(this),
      error: function(err){
          alert('Error adding article');
      }
    });
  },
  render: function() {
    var listElts = [];
    for(var i = 0; i < this.state.articleTitles.length; i++) {
      var title = this.state.articleTitles[i];
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
              <input
                name="formArticleTitle"
                type="text"
                className="form-control"
                value={this.state.formArticleTitle}
                id="articleTitle"
                placeholder="Title"
                onChange={this.handleChange}/>
            </div>
            <div className="form-group">
              <label for="articleTitle">Article Content</label>
              <input
                name="formArticleContent"
                type="text"
                className="form-control"
                value={this.state.formArticleContent}
                id="articleContent"
                placeholder="Content"
                onChange={this.handleChange}/>
            </div>
            <button className="btn btn-default">Create</button>
          </form>
        </div>
      </div>
    );
  },
})

ReactDOM.render(<Root/>, document.getElementById("page-content"));
