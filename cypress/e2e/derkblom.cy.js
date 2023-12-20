describe('Home page', () => {
  it('tests the dark/light mode button', () => {
    cy.visit('http://localhost:5173/')

    cy.wait(1000)

    cy.get(".lightswitch-track").scrollIntoView()
    cy.get(".lightswitch-track").should('be.visible');
    cy.get(".lightswitch-track").trigger("click")
    cy.get("body").should("have.css", "background-color", "rgb(36, 44, 70)")
    // div lightswitch-thumb
    
  })
})