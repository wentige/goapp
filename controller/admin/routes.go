package admin

import echo "github.com/labstack/echo/v4"

func RegisterRoutes(g *echo.Group) {
	g.GET("", AdminIndex)
	new(AuthorityController).RegisterRoute(g)
	new(UserController).RegisterRoute(g)
	new(TopicController).RegisterRoute(g)
	new(NodeController).RegisterRoute(g)
	new(ArticleController).RegisterRoute(g)
	new(ProjectController).RegisterRoute(g)
	new(RuleController).RegisterRoute(g)
	new(ReadingController).RegisterRoute(g)
	new(ToolController).RegisterRoute(g)
	new(SettingController).RegisterRoute(g)
	new(MetricsController).RegisterRoute(g)
}
