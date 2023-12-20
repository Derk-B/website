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
})