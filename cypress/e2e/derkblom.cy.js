describe('Home page', () => {
  it('tests the dark/light mode button', () => {
    cy.visit('http://localhost:5173/')

    cy.wait(500)

    cy.get(".lightswitch-track").scrollIntoView()
    cy.get(".lightswitch-track").should('be.visible');
    cy.get(".lightswitch-track").trigger("click")
    cy.get("body").should("have.css", "background-color", "rgb(36, 44, 70)")
  })
})

describe('Projects page', () => {
  it('tests if navigation to projects page works', () => {
    cy.visit('http://localhost:5173')
    cy.get('a[href="/projects"]').click()
    cy.location().should((location) => {
      expect(location.href).to.eq("http://localhost:5173/projects")
    })
  })

  it('tests if ate least 1 project cards are displayed', () => {
    cy.visit('http://localhost:5173/projects')
    cy.get('main').find('a').should('have.length.at.least', 1)
  })

  it('tests if the project card contains all necessary elements', () => {
    cy.visit('http://localhost:5173/projects')
    cy.get('main').find('a').each(($item, index, $itemList) => {
      cy.wrap($item).find('img').should("exist")
      cy.wrap($item).find('h2').should("exist")
      cy.wrap($item).find('p').should("exist")
    })
  })
})

describe('Blog page', () => {
  it('tests if navigation to blog page works', () => {
    cy.visit('http://localhost:5173')
    cy.get('a[href="/blog"]').click()
    cy.location().should((location) => {
      expect(location.href).to.eq("http://localhost:5173/blog")
    })
  })

  it('tests if ate least 1 blog cards are displayed', () => {
    cy.visit('http://localhost:5173/blog')
    cy.get('main').find('a').should('have.length.at.least', 1)
  })

  it('tests if the project card contains all necessary elements', () => {
    cy.visit('http://localhost:5173/blog')
    cy.get('main').find('a').each(($item, index, $itemList) => {
      cy.wrap($item).find('img').should("exist")
      cy.wrap($item).find('h2').should("exist")
      cy.wrap($item).find('p').should("exist")
    })
  })

  it('tests if the user is sent to the correct blog page', () => {
    cy.visit('http://localhost:5173/blog')
    cy.get('main').find('a').first().click()

    cy.location("href").should('include', "http://localhost:5173/blogpost")
  })
})

describe('blogpost page', () => {
  it("tests if the page contains an image", () => {
    cy.visit('http://localhost:5173/blog')
    cy.get('main').find('a').first().click()

    cy.get('h1').should('exist')
    cy.get('img').should('exist')
  })
})