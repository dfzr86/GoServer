package article

import (
	"comment"
	"database/sql"
	"fmt"
	"mysql"
)

var dbName = "t_article"

type Article struct {
	Id					string `json:"id"`
	Title				string `json:"title"`
	TitleDesc			string `json:"desc"`
	CoverImage			string `json:"cover_image"`//封面图片
	ContentPath			string `json:"content_path"`//音频文件路径
}

func GetArticleList(size int, cursor string) ([]Article, error) {

	//去db根据cursor查找指定的条数,

	//变成json返回去
	return []Article{}, nil
}

func (a *Article)getComments() ([]comment.Comment, error) {

	//根据ArticleId去获取对应的comment
	return nil, nil

}

//获取文章
func Get(aid string) (*Article, error)  {

	stmt, err := mysql.SharedDb().Prepare("SELECT id, title, titleDesc, coverImage, contentPath FROM t_article WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	a := Article{}

	articleId := sql.NullString{}
	title := sql.NullString{}
	desc := sql.NullString{}
	coverImage := sql.NullString{}
	contentPath := sql.NullString{}

	err = stmt.QueryRow(aid).Scan(&articleId, &title, &desc, &coverImage, &contentPath)
	if err != nil {
		panic(err.Error())
		return nil, err
	}

	a.Id = articleId.String
	a.Title = title.String
	a.TitleDesc = desc.String
	a.CoverImage = coverImage.String
	a.ContentPath = contentPath.String

	fmt.Println(articleId, title, desc, coverImage, contentPath)

	return &a, nil
}

//删除文章
func Delete(id string) error  {

	_, err := mysql.SharedDb().Exec("DELETE FROM t_article WHERE id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
		return err;
	}

	return nil;
}

//修改文章


//新增文章
func (a *Article) Create() error {

	//stmt, err := sql.SharedDb.Prepare("INSERT INTO ?")
	return nil

}

func (a *Article) Modify() error {
	return nil


}





