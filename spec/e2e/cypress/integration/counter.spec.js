describe('/counter', () => {
  beforeEach(() => {
    cy.visit(Cypress.env('DEV_TEST_URL') + '/counter')
  })

  it('increments counter and resets', () => {
    // three counter
    cy.get('#counter_content > button:nth-child(4)')
      .contains('+')
      .click()
      .click()
      .click()
    cy.get('#counter_content > p').should('have.text', '3')

    // thirty counter
    cy.get('#counter_content > button:nth-child(5)')
      .contains('+10')
      .click()
      .click()
      .click()
    cy.get('#counter_content > p').should('have.text', '33')
  })

  it('decrements counter and resets', () => {
    // three counter
    cy.get('#counter_content > button:nth-child(2)')
      .contains('-')
      .click()
      .click()
      .click()
    cy.get('#counter_content > p').should('have.text', '-3')

    // thirty counter
    cy.get('#counter_content > button:nth-child(1)')
      .contains('-10')
      .click()
      .click()
      .click()
    cy.get('#counter_content > p').should('have.text', '-33')
  })

  it('resets counter', () => {
    // thirty counter before resets
    cy.get('#counter_content > button')
      .contains('+10')
      .click()
      .click()
      .click()

    cy.get('#reset').click()
    cy.get('#counter_content > p').should('have.text', '0')
  })
})