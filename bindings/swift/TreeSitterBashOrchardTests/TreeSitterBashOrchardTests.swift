import XCTest
import SwiftTreeSitter
import TreeSitterBashOrchard

final class TreeSitterBashOrchardTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_bash_orchard())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading Bash grammar")
    }
}
