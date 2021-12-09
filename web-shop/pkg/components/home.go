package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type MyComponent struct {
	app.Compo
}

func (c *MyComponent) Render() app.UI {
	return app.Div().
		Class("pf-c-page").
		ID("card-view-example").
		Body(
			app.A().
				Class("pf-c-skip-to-content pf-c-button pf-m-primary").
				Href("#main-content-card-view-example").
				Body(
					app.Text("Skip to content"),
				),
			app.Header().
				Class("pf-c-page__header").
				Body(
					app.Div().
						Class("pf-c-page__header-brand").
						Body(
							app.Div().
								Class("pf-c-page__header-brand-toggle").
								Body(
									app.Button().
										Class("pf-c-button pf-m-plain").
										Type("button").
										ID("card-view-example-nav-toggle").
										Aria("label", "Global navigation").
										Aria("expanded", true).
										Aria("controls", "card-view-example-primary-nav").
										Body(
											app.I().
												Class("fas fa-bars").
												Aria("hidden", true),
										),
								),
							app.A().
								Href("#").
								Class("pf-c-page__header-brand-link").
								Body(
									app.Img().
										Class("pf-c-brand").
										Src("https://flash---art.com/wp-content/uploads/2020/08/Tom%C3%A1s-Saraceno_Arachnid-Orchestra_331_Flash-Art-05.jpg").
										Alt("Web shop logo"),
								),
						),
					app.Div().
						Class("pf-c-page__header-tools").
						Body(
							app.Div().
								Class("pf-c-page__header-tools-group").
								Body(
									app.Div().
										Class("pf-c-page__header-tools-item pf-m-hidden pf-m-visible-on-lg").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Settings").
												Body(
													app.I().
														Class("fas fa-cog").
														Aria("hidden", true),
												),
										),
									app.Div().
										Class("pf-c-page__header-tools-item pf-m-hidden pf-m-visible-on-lg").
										Body(
											app.Button().
												Class("pf-c-button pf-m-plain").
												Type("button").
												Aria("label", "Help").
												Body(
													app.I().
														Class("pf-icon pf-icon-help").
														Aria("hidden", true),
												),
										),
								),
							app.Div().
								Class("pf-c-page__header-tools-group").
								Body(
									app.Div().
										Class("pf-c-page__header-tools-item pf-m-hidden-on-lg").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("card-view-example-dropdown-kebab-1-button").
														Aria("expanded", "false").
														Type("button").
														Aria("label", "Actions").
														Body(
															app.I().
																Class("fas fa-ellipsis-v").
																Aria("hidden", true),
														),
													app.Ul().
														Class("pf-c-dropdown__menu pf-m-align-right").
														Aria("labelledby", "card-view-example-dropdown-kebab-1-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
																		Href("#").
																		Aria("disabled", true).
																		TabIndex(-1).
																		Body(
																			app.Text("Disabled link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
									app.Div().
										Class("pf-c-page__header-tools-item pf-m-hidden pf-m-visible-on-md").
										Body(
											app.Div().
												Class("pf-c-dropdown").
												Body(
													app.Button().
														Class("pf-c-dropdown__toggle pf-m-plain").
														ID("card-view-example-dropdown-kebab-2-button").
														Aria("expanded", "false").
														Type("button").
														Body(
															app.Span().
																Class("pf-c-dropdown__toggle-text").
																Body(
																	app.Text("Account"),
																),
															app.Span().
																Class("pf-c-dropdown__toggle-icon").
																Body(
																	app.I().
																		Class("fas fa-caret-down").
																		Aria("hidden", true),
																),
														),
													app.Ul().
														Class("pf-c-dropdown__menu").
														Aria("labelledby", "card-view-example-dropdown-kebab-2-button").
														Hidden(true).
														Body(
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Body(
																			app.Text("Action"),
																		),
																),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item pf-m-disabled").
																		Href("#").
																		Aria("disabled", true).
																		TabIndex(-1).
																		Body(
																			app.Text("Disabled link"),
																		),
																),
															app.Li().
																Body(
																	app.Button().
																		Class("pf-c-dropdown__menu-item").
																		Type("button").
																		Disabled(true).
																		Body(
																			app.Text("Disabled action"),
																		),
																),
															app.Li().
																Class("pf-c-divider").
																Aria("role", "separator"),
															app.Li().
																Body(
																	app.A().
																		Class("pf-c-dropdown__menu-item").
																		Href("#").
																		Body(
																			app.Text("Separated link"),
																		),
																),
														),
												),
										),
								),
							app.Img().
								Class("pf-c-avatar").
								Src("https://www.lego.com/cdn/cs/set/assets/blt13e6e2a178c38704/Spiderman-Sidekick-Tall-1.jpg?fit=crop&format=jpg&quality=80&width=800&height=600&dpr=1").
								Alt("Avatar image"),
						),
				),
			app.Div().
				Class("pf-c-page__sidebar").
				Body(
					app.Div().
						Class("pf-c-page__sidebar-body").
						Body(
							app.Nav().
								Class("pf-c-nav").
								ID("card-view-example-primary-nav").
								Aria("label", "Global").
								Body(
									app.Ul().
										Class("pf-c-nav__list").
										Body(
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link pf-m-current").
														Aria("current", "page").
														Body(
															app.Text("Spiral orb webs"),
														),
												),
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link").
														Body(
															app.Text("Tangle webs"),
														),
												),
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link").
														Body(
															app.Text("Funnel webs"),
														),
												),
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link").
														Body(
															app.Text("Turbular webs"),
														),
												),
											app.Li().
												Class("pf-c-nav__item").
												Body(
													app.A().
														Href("#").
														Class("pf-c-nav__link").
														Body(
															app.Text("Sheet webs"),
														),
												),
										),
								),
						),
				),
			app.Main().
				Class("pf-c-page__main").
				TabIndex(-1).
				ID("main-content-card-view-example").
				Body(
					app.Section().
						Class("pf-c-page__main-section pf-m-light").
						Body(
							app.Div().
								Class("pf-c-content").
								Body(
									app.H1().
										Body(
											app.Text("Webs"),
										),
									app.P().
										Body(
											app.Text("Buy the best quality webs straight outta Internet"),
										),
								),
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-light pf-m-no-padding").
						Body(
							app.Div().
								Class("pf-c-toolbar").
								Body(
									app.Div().
										Class("pf-c-toolbar__content").
										Body(
											app.Div().
												Class("pf-c-toolbar__content-section pf-m-nowrap").
												Body(
													app.Div().
														Class("pf-c-toolbar__group pf-m-toggle-group pf-m-show-on-xl").
														Body(
															app.Div().
																Class("pf-c-toolbar__toggle").
																Body(
																	app.Button().
																		Class("pf-c-button pf-m-plain").
																		Type("button").
																		Aria("label", "Show filters").
																		Aria("expanded", "false").
																		Aria("controls", "-expandable-content").
																		Body(
																			app.I().
																				Class("fas fa-filter").
																				Aria("hidden", true),
																		),
																),
															app.Div().
																Class("pf-c-toolbar__item pf-m-bulk-select").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Div().
																				Class("pf-c-dropdown__toggle pf-m-split-button").
																				Body(
																					app.Label().
																						Class("pf-c-dropdown__toggle-check").
																						For("-bulk-select-toggle-check").
																						Body(
																							app.Input().
																								Type("checkbox").
																								ID("-bulk-select-toggle-check").
																								Aria("label", "Select all"),
																						),
																					app.Button().
																						Class("pf-c-dropdown__toggle-button").
																						Type("button").
																						Aria("expanded", "false").
																						ID("-bulk-select-toggle-button").
																						Aria("label", "Dropdown toggle").
																						Body(
																							app.I().
																								Class("fas fa-caret-down").
																								Aria("hidden", true),
																						),
																				),
																			app.Ul().
																				Class("pf-c-dropdown__menu").
																				Hidden(true).
																				Body(
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-dropdown__menu-item").
																								Type("button").
																								Body(
																									app.Text("Select all"),
																								),
																						),
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-dropdown__menu-item").
																								Type("button").
																								Body(
																									app.Text("Select none"),
																								),
																						),
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-dropdown__menu-item").
																								Type("button").
																								Body(
																									app.Text("Other action"),
																								),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-toolbar__item").
																Body(
																	app.Div().
																		Class("pf-c-select").
																		Body(
																			app.Span().
																				ID("-select-checkbox-status-label").
																				Hidden(true).
																				Body(
																					app.Text("Choose one"),
																				),
																			app.Button().
																				Class("pf-c-select__toggle").
																				Type("button").
																				ID("-select-checkbox-status-toggle").
																				Aria("haspopup", true).
																				Aria("expanded", "false").
																				Aria("labelledby", "-select-checkbox-status-label -select-checkbox-status-toggle").
																				Body(
																					app.Div().
																						Class("pf-c-select__toggle-wrapper").
																						Body(
																							app.Span().
																								Class("pf-c-select__toggle-text").
																								Body(
																									app.Text("Spider"),
																								),
																						),
																					app.Span().
																						Class("pf-c-select__toggle-arrow").
																						Body(
																							app.I().
																								Class("fas fa-caret-down").
																								Aria("hidden", true),
																						),
																				),
																			app.Div().
																				Class("pf-c-select__menu").
																				Hidden(true).
																				Body(
																					app.FieldSet().
																						Class("pf-c-select__menu-fieldset").
																						Aria("label", "Select input").
																						Body(
																							app.Label().
																								Class("pf-c-check pf-c-select__menu-item pf-m-description").
																								For("-select-checkbox-status-active").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("-select-checkbox-status-active").
																										Name("-select-checkbox-status-active"),
																									app.Span().
																										Class("pf-c-check__label").
																										Body(
																											app.Text("Active"),
																										),
																									app.Span().
																										Class("pf-c-check__description").
																										Body(
																											app.Text("This is a description"),
																										),
																								),
																							app.Label().
																								Class("pf-c-check pf-c-select__menu-item pf-m-description").
																								For("-select-checkbox-status-canceled").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("-select-checkbox-status-canceled").
																										Name("-select-checkbox-status-canceled"),
																									app.Span().
																										Class("pf-c-check__label").
																										Body(
																											app.Text("Canceled"),
																										),
																									app.Span().
																										Class("pf-c-check__description").
																										Body(
																											app.Text("This is a really long description that describes the menu item. This is a really long description that describes the menu item."),
																										),
																								),
																							app.Label().
																								Class("pf-c-check pf-c-select__menu-item").
																								For("-select-checkbox-status-paused").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("-select-checkbox-status-paused").
																										Name("-select-checkbox-status-paused"),
																									app.Span().
																										Class("pf-c-check__label").
																										Body(
																											app.Text("Paused"),
																										),
																								),
																							app.Label().
																								Class("pf-c-check pf-c-select__menu-item").
																								For("-select-checkbox-status-warning").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("-select-checkbox-status-warning").
																										Name("-select-checkbox-status-warning"),
																									app.Span().
																										Class("pf-c-check__label").
																										Body(
																											app.Text("Warning"),
																										),
																								),
																							app.Label().
																								Class("pf-c-check pf-c-select__menu-item").
																								For("-select-checkbox-status-restarted").
																								Body(
																									app.Input().
																										Class("pf-c-check__input").
																										Type("checkbox").
																										ID("-select-checkbox-status-restarted").
																										Name("-select-checkbox-status-restarted"),
																									app.Span().
																										Class("pf-c-check__label").
																										Body(
																											app.Text("Restarted"),
																										),
																								),
																						),
																				),
																		),
																),
														),
													app.Div().
														Class("pf-c-overflow-menu").
														ID("-overflow-menu").
														Body(

															app.Div().
																Class("pf-c-overflow-menu__control").
																Body(
																	app.Div().
																		Class("pf-c-dropdown").
																		Body(
																			app.Button().
																				Class("pf-c-button pf-c-dropdown__toggle pf-m-plain").
																				Type("button").
																				ID("-overflow-menu-dropdown-toggle").
																				Aria("label", "Dropdown with additional options").
																				Aria("expanded", "false").
																				Body(
																					app.I().
																						Class("fas fa-ellipsis-v").
																						Aria("hidden", true),
																				),
																			app.Ul().
																				Class("pf-c-dropdown__menu").
																				Aria("labelledby", "-overflow-menu-dropdown-toggle").
																				Hidden(true).
																				Body(
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-dropdown__menu-item").
																								Body(
																									app.Text("Action 7"),
																								),
																						),
																				),
																		),
																),
														),
													app.Div().
														Class("pf-c-toolbar__item pf-m-pagination").
														Body(
															app.Div().
																Class("pf-c-pagination pf-m-compact").
																Body(
																	app.Div().
																		Class("pf-c-options-menu").
																		Body(
																			app.Div().
																				Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
																				Body(
																					app.Span().
																						Class("pf-c-options-menu__toggle-text").
																						Body(
																							app.B().
																								Body(
																									app.Text("1 - 10"),
																								),
																							app.Text("of"),
																							app.B().
																								Body(
																									app.Text("36"),
																								),
																						),
																					app.Button().
																						Class("pf-c-options-menu__toggle-button").
																						ID("-top-pagination-toggle").
																						Aria("haspopup", "listbox").
																						Aria("expanded", "false").
																						Aria("label", "Items per page").
																						Body(
																							app.Span().
																								Class("pf-c-options-menu__toggle-button-icon").
																								Body(
																									app.I().
																										Class("fas fa-caret-down").
																										Aria("hidden", true),
																								),
																						),
																				),
																			app.Ul().
																				Class("pf-c-options-menu__menu").
																				Aria("labelledby", "-top-pagination-toggle").
																				Hidden(true).
																				Body(
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-options-menu__menu-item").
																								Type("button").
																								Body(
																									app.Text("5 per page"),
																								),
																						),
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-options-menu__menu-item").
																								Type("button").
																								Body(
																									app.Text("10 per page"), app.Div().
																										Class("pf-c-options-menu__menu-item-icon").
																										Body(
																											app.I().
																												Class("fas fa-check").
																												Aria("hidden", true),
																										),
																								),
																						),
																					app.Li().
																						Body(
																							app.Button().
																								Class("pf-c-options-menu__menu-item").
																								Type("button").
																								Body(
																									app.Text("20 per page"),
																								),
																						),
																				),
																		),
																	app.Nav().
																		Class("pf-c-pagination__nav").
																		Aria("label", "Toolbar top pagination").
																		Body(
																			app.Div().
																				Class("pf-c-pagination__nav-control pf-m-prev").
																				Body(
																					app.Button().
																						Class("pf-c-button pf-m-plain").
																						Type("button").
																						Disabled(true).
																						Aria("label", "Go to previous page").
																						Body(
																							app.I().
																								Class("fas fa-angle-left").
																								Aria("hidden", true),
																						),
																				),
																			app.Div().
																				Class("pf-c-pagination__nav-control pf-m-next").
																				Body(
																					app.Button().
																						Class("pf-c-button pf-m-plain").
																						Type("button").
																						Aria("label", "Go to next page").
																						Body(
																							app.I().
																								Class("fas fa-angle-right").
																								Aria("hidden", true),
																						),
																				),
																		),
																),
														),
												),
											app.Div().
												Class("pf-c-toolbar__expandable-content pf-m-hidden").
												ID("-expandable-content").
												Hidden(true),
										),
								),
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-fill").
						Body(
							app.Div().
								Class("pf-l-gallery pf-m-gutter").
								Body(
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.thetimes.co.uk/imageserver/image/%2Fmethode%2Ftimes%2Fprod%2Fweb%2Fbin%2F57001458-e868-11e9-b931-c019e957f02a.jpg?crop=4767%2C2681%2C0%2C248&resize=1200").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://singularityhub.com/wp-content/uploads/2019/03/spider-web-silk-close-dark-biotechnology-shutterstock-200612795-1068x601.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.reuters.com/resizer/mkf6qE8fOHLiqSBt3N0lRPwLass=/960x0/filters:quality(80)/cloudfront-us-east-2.images.arcpublishing.com/reuters/XAZDQDPV3VJRVJH2CCQBJRJ7EI.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.purdue.edu/uns/images/2020/lee-websLO.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.thetimes.co.uk/imageserver/image/%2Fmethode%2Ftimes%2Fprod%2Fweb%2Fbin%2F57001458-e868-11e9-b931-c019e957f02a.jpg?crop=4767%2C2681%2C0%2C248&resize=1200").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://flash---art.com/wp-content/uploads/2020/08/Tom%C3%A1s-Saraceno_Arachnid-Orchestra_331_Flash-Art-05.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.purdue.edu/uns/images/2020/lee-websLO.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://im.indiatimes.in/content/2021/Apr/spider-web_607563624207c.jpeg?w=725&h=407").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://singularityhub.com/wp-content/uploads/2019/03/spider-web-silk-close-dark-biotechnology-shutterstock-200612795-1068x601.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://im.indiatimes.in/content/2021/Apr/spider-web_607563624207c.jpeg?w=725&h=407").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.jordantimes.com/sites/default/files/styles/news_inner/public/focus_48.jpg?itok=PgJ_EuwP").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.reuters.com/resizer/mkf6qE8fOHLiqSBt3N0lRPwLass=/960x0/filters:quality(80)/cloudfront-us-east-2.images.arcpublishing.com/reuters/XAZDQDPV3VJRVJH2CCQBJRJ7EI.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.purdue.edu/uns/images/2020/lee-websLO.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.reuters.com/resizer/mkf6qE8fOHLiqSBt3N0lRPwLass=/960x0/filters:quality(80)/cloudfront-us-east-2.images.arcpublishing.com/reuters/XAZDQDPV3VJRVJH2CCQBJRJ7EI.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://im.indiatimes.in/content/2021/Apr/spider-web_607563624207c.jpeg?w=725&h=407").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.loredohands.com/images/spider-web-trap-for-insects.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.jordantimes.com/sites/default/files/styles/news_inner/public/focus_48.jpg?itok=PgJ_EuwP").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.loredohands.com/images/spider-web-trap-for-insects.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.thetimes.co.uk/imageserver/image/%2Fmethode%2Ftimes%2Fprod%2Fweb%2Fbin%2F57001458-e868-11e9-b931-c019e957f02a.jpg?crop=4767%2C2681%2C0%2C248&resize=1200").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.loredohands.com/images/spider-web-trap-for-insects.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://singularityhub.com/wp-content/uploads/2019/03/spider-web-silk-close-dark-biotechnology-shutterstock-200612795-1068x601.jpg").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),

									app.Div().
										Class("pf-c-card pf-m-selectable-raised pf-m-compact").
										ID("card-1").
										Body(
											app.Div().
												Class("pf-c-card__header").
												Body(
													app.Img().
														Src("https://www.jordantimes.com/sites/default/files/styles/news_inner/public/focus_48.jpg?itok=PgJ_EuwP").
														Alt("Spider web logo"),
													app.Div().
														Class("pf-c-card__actions").
														Body(
															app.Div().
																Class("pf-c-dropdown").
																Body(
																	app.Button().
																		Class("pf-c-dropdown__toggle pf-m-plain").
																		ID("card-1-dropdown-kebab-button").
																		Aria("expanded", "false").
																		Type("button").
																		Aria("label", "Actions").
																		Body(
																			app.I().
																				Class("fas fa-ellipsis-v").
																				Aria("hidden", true),
																		),
																	app.Ul().
																		Class("pf-c-dropdown__menu pf-m-align-right").
																		Aria("labelledby", "card-1-dropdown-kebab-button").
																		Hidden(true).
																		Body(
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Body(
																							app.Text("Action"),
																						),
																				),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item pf-m-disabled").
																						Href("#").
																						Aria("disabled", true).
																						TabIndex(-1).
																						Body(
																							app.Text("Disabled link"),
																						),
																				),
																			app.Li().
																				Body(
																					app.Button().
																						Class("pf-c-dropdown__menu-item").
																						Type("button").
																						Disabled(true).
																						Body(
																							app.Text("Disabled action"),
																						),
																				),
																			app.Li().
																				Class("pf-c-divider").
																				Aria("role", "separator"),
																			app.Li().
																				Body(
																					app.A().
																						Class("pf-c-dropdown__menu-item").
																						Href("#").
																						Body(
																							app.Text("Separated link"),
																						),
																				),
																		),
																),
															app.Div().
																Class("pf-c-check pf-m-standalone").
																Body(
																	app.Input().
																		Class("pf-c-check__input").
																		Type("checkbox").
																		ID("card-1-check").
																		Name("card-1-check").
																		Aria("labelledby", "card-1-check-label"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__title").
												Body(
													app.P().
														ID("card-1-check-label").
														Body(
															app.Text("Spider web"),
														),
													app.Div().
														Class("pf-c-content").
														Body(
															app.Small().
																Body(
																	app.Text("Spider not included"),
																),
														),
												),
											app.Div().
												Class("pf-c-card__body").
												Body(
													app.Text("Spider web is a SPIDER PIG SPIDER PIG Does whatever a SPIDER PIG does Can he swingFrom a web No he cant He's a pig LOOK OOOUUUTTT!!!! He is a SPIDER PIG!! that promotes design commonality and improves user experience."),
												),
										),
								),
						),
					app.Section().
						Class("pf-c-page__main-section pf-m-no-padding pf-m-light pf-m-sticky-bottom pf-m-no-fill").
						Body(
							app.Div().
								Class("pf-c-pagination pf-m-bottom").
								Body(
									app.Div().
										Class("pf-c-options-menu pf-m-top").
										Body(
											app.Div().
												Class("pf-c-options-menu__toggle pf-m-text pf-m-plain").
												Body(
													app.Span().
														Class("pf-c-options-menu__toggle-text").
														Body(
															app.B().
																Body(
																	app.Text("1 - 10"),
																),
															app.Text("of"),
															app.B().
																Body(
																	app.Text("36"),
																),
														),
													app.Button().
														Class("pf-c-options-menu__toggle-button").
														ID("pagination-options-menu-bottom-example-toggle").
														Aria("haspopup", "listbox").
														Aria("expanded", "false").
														Aria("label", "Items per page").
														Body(
															app.Span().
																Class("pf-c-options-menu__toggle-button-icon").
																Body(
																	app.I().
																		Class("fas fa-caret-down").
																		Aria("hidden", true),
																),
														),
												),
											app.Ul().
												Class("pf-c-options-menu__menu pf-m-top").
												Aria("labelledby", "pagination-options-menu-bottom-example-toggle").
												Hidden(true).
												Body(
													app.Li().
														Body(
															app.Button().
																Class("pf-c-options-menu__menu-item").
																Type("button").
																Body(
																	app.Text("5 per page"),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-options-menu__menu-item").
																Type("button").
																Body(
																	app.Text("10 per page"), app.Div().
																		Class("pf-c-options-menu__menu-item-icon").
																		Body(
																			app.I().
																				Class("fas fa-check").
																				Aria("hidden", true),
																		),
																),
														),
													app.Li().
														Body(
															app.Button().
																Class("pf-c-options-menu__menu-item").
																Type("button").
																Body(
																	app.Text("20 per page"),
																),
														),
												),
										),
									app.Nav().
										Class("pf-c-pagination__nav").
										Aria("label", "Pagination").
										Body(
											app.Div().
												Class("pf-c-pagination__nav-control pf-m-first").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Disabled(true).
														Aria("label", "Go to first page").
														Body(
															app.I().
																Class("fas fa-angle-double-left").
																Aria("hidden", true),
														),
												),
											app.Div().
												Class("pf-c-pagination__nav-control pf-m-prev").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Disabled(true).
														Aria("label", "Go to previous page").
														Body(
															app.I().
																Class("fas fa-angle-left").
																Aria("hidden", true),
														),
												),
											app.Div().
												Class("pf-c-pagination__nav-page-select").
												Body(
													app.Input().
														Class("pf-c-form-control").
														Aria("label", "Current page").
														Type("number").
														Min("1").
														Max("4").
														Value("1"),
													app.Span().
														Aria("hidden", true).
														Body(
															app.Text("of 4"),
														),
												),
											app.Div().
												Class("pf-c-pagination__nav-control pf-m-next").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Go to next page").
														Body(
															app.I().
																Class("fas fa-angle-right").
																Aria("hidden", true),
														),
												),
											app.Div().
												Class("pf-c-pagination__nav-control pf-m-last").
												Body(
													app.Button().
														Class("pf-c-button pf-m-plain").
														Type("button").
														Aria("label", "Go to last page").
														Body(
															app.I().
																Class("fas fa-angle-double-right").
																Aria("hidden", true),
														),
												),
										),
								),
						),
				),
		)
}
